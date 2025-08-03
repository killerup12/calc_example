package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"calc_example/internal/config"
	"calc_example/internal/handler"
	"calc_example/internal/repository"
	"calc_example/internal/service"
	"calc_example/pkg/database"
	"calc_example/pkg/logger"

	"github.com/gin-gonic/gin"
)

type App struct {
	config  *config.Config
	logger  *logger.Logger
	router  *gin.Engine
	server  *http.Server
	db      *database.Database
	handler *handler.Handler
	service *service.Service
	repo    *repository.Repository
}

func New(cfg *config.Config) *App {
	// Инициализируем логгер
	log := logger.New(cfg.Log.Level)

	// Инициализируем базу данных
	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Инициализируем репозиторий
	repo := repository.New(db)

	// Инициализируем сервисы
	services := service.New(repo)

	// Инициализируем хендлеры
	handlers := handler.New(services, log)

	// Инициализируем роутер
	router := gin.Default()

	// Настраиваем middleware
	setupMiddleware(router, log)

	// Настраиваем роуты
	handlers.InitRoutes(router)

	// Создаем HTTP сервер
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	return &App{
		config:  cfg,
		logger:  log,
		router:  router,
		server:  server,
		db:      db,
		handler: handlers,
		service: services,
		repo:    repo,
	}
}

func (a *App) Run() error {
	// Запускаем сервер в горутине
	go func() {
		a.logger.Info("Сервер запущен на:", a.config.Server.Host, ":", a.config.Server.Port)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal("Ошибка запуска сервера:", err)
		}
	}()

	// Ждем сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.logger.Info("Получен сигнал завершения, закрываем сервер...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Fatal("Ошибка при закрытии сервера:", err)
	}

	// Закрываем соединение с базой данных
	if err := a.db.Close(); err != nil {
		a.logger.Error("Ошибка при закрытии базы данных:", err)
	}

	a.logger.Info("Сервер успешно остановлен")
	return nil
}

func setupMiddleware(router *gin.Engine, log *logger.Logger) {
	// Логирование запросов
	router.Use(gin.Logger())

	// Recovery middleware
	router.Use(gin.Recovery())

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

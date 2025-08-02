package main

import (
	"log"

	"calc_example/internal/app"
	"calc_example/internal/config"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Создаем и запускаем приложение
	application := app.New(cfg)
	
	if err := application.Run(); err != nil {
		log.Fatalf("Ошибка запуска приложения: %v", err)
	}
} 
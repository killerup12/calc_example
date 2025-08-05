package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"calc_example/internal/model"
	"calc_example/internal/service"
	"calc_example/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	logger  *logger.Logger
}

func New(service *service.Service, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	// Группа API v1
	api := router.Group("/api/v1")
	{
		// Заявки
		api.POST("/issue", h.createIssue)
		api.GET("/issues", h.getAllIssues)
		api.GET("/issue/:id", h.getIssueByID)
		api.PATCH("/issue/:id", h.updateIssue)
	}

	// Health check
	router.GET("/health", h.healthCheck)
}

// Issue handlers
func (h *Handler) createIssue(c *gin.Context) {
	var req model.CreateIssueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Ошибка валидации запроса:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	issue, err := h.service.CreateIssue(&req)
	if err != nil {
		h.logger.Error("Ошибка создания заявки:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Отправка сообщения в Telegram
	err = sendTelegramMessage(issue) // Передаем req для формирования сообщения
	if err != nil {
		h.logger.Error("Ошибка отправки сообщения в Telegram:", err)
	}

	c.JSON(http.StatusCreated, issue)
}

func sendTelegramMessage(issue *model.IssueResponse) error {
	url := os.Getenv("TELEGRAM_BOT_SERVICE") + "/send-message"

	data := map[string]string{
		"text": fmt.Sprintf("🆕<b>Новая заявка</b>:\n\n👤 Имя: %s\n📞Телефон: %s\n\n📦 Товар: %s\n📲 Источник: %s\n\n🧑🏻‍💻Менеджер: %s\n📌 Статус: %s",
			issue.FullName, issue.ContactInfo, issue.ProductDescription, "Сайт", "Виртуальный помощник", "Ожидает ответа"),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("_ошибка при преобразовании в JSON: %v", err)
	}

	client := &http.Client{}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("__ошибка при создании запроса: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("___ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Читаем тело ответа для диагностики
		return fmt.Errorf("неуспешный статус ответа: %s, тело: %s", resp.Status, string(body))
	}

	return nil
}

func (h *Handler) getAllIssues(c *gin.Context) {
	issues, err := h.service.GetAllIssues()
	if err != nil {
		h.logger.Error("Ошибка получения заявок:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	c.JSON(http.StatusOK, issues)
}

func (h *Handler) getIssueByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID заявки"})
		return
	}

	issue, err := h.service.GetIssueByID(uint(id))
	if err != nil {
		h.logger.Error("Ошибка получения заявки:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	c.JSON(http.StatusOK, issue)
}

func (h *Handler) updateIssue(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID заявки"})
		return
	}

	var req model.UpdateIssueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Ошибка валидации запроса:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	issue, err := h.service.UpdateIssue(uint(id), &req)
	if err != nil {
		h.logger.Error("Ошибка обновления заявки:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, issue)
}

// Health check
func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Сервер работает",
	})
}

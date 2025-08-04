package handler

import (
	"net/http"
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

	c.JSON(http.StatusCreated, issue)
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

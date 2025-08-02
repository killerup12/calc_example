.PHONY: build run test clean help

# Переменные
BINARY_NAME=calc_example
BUILD_DIR=bin
MAIN_FILE=cmd/server/main.go

# Цвета для вывода
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

help: ## Показать справку
	@echo "$(GREEN)Доступные команды:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

build: ## Собрать приложение
	@echo "$(GREEN)Сборка приложения...$(NC)"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "$(GREEN)Приложение собрано: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

run: ## Запустить сервер в режиме разработки
	@echo "$(GREEN)Запуск сервера...$(NC)"
	go run $(MAIN_FILE)

test: ## Запустить тесты
	@echo "$(GREEN)Запуск тестов...$(NC)"
	go test -v ./...

test-coverage: ## Запустить тесты с покрытием
	@echo "$(GREEN)Запуск тестов с покрытием...$(NC)"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Отчет о покрытии сохранен в coverage.html$(NC)"

clean: ## Очистить собранные файлы
	@echo "$(GREEN)Очистка...$(NC)"
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	rm -f *.db

deps: ## Установить зависимости
	@echo "$(GREEN)Установка зависимостей...$(NC)"
	go mod tidy
	go mod download

fmt: ## Форматировать код
	@echo "$(GREEN)Форматирование кода...$(NC)"
	go fmt ./...

lint: ## Проверить код линтером
	@echo "$(GREEN)Проверка кода...$(NC)"
	golangci-lint run

dev: ## Запустить в режиме разработки с автоперезагрузкой
	@echo "$(GREEN)Запуск в режиме разработки...$(NC)"
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "$(YELLOW)Air не установлен. Установите: go install github.com/cosmtrek/air@latest$(NC)"; \
		go run $(MAIN_FILE); \
	fi

install-air: ## Установить Air для автоперезагрузки
	@echo "$(GREEN)Установка Air...$(NC)"
	go install github.com/cosmtrek/air@latest

docker-build: ## Собрать Docker образ
	@echo "$(GREEN)Сборка Docker образа...$(NC)"
	docker build -t $(BINARY_NAME) .

docker-run: ## Запустить в Docker
	@echo "$(GREEN)Запуск в Docker...$(NC)"
	docker run -p 8080:8080 $(BINARY_NAME)

# Команды для работы с базой данных
db-migrate: ## Выполнить миграции БД
	@echo "$(GREEN)Миграции выполняются автоматически при запуске$(NC)"

db-reset: ## Сбросить базу данных
	@echo "$(GREEN)Сброс базы данных...$(NC)"
	rm -f *.db
	@echo "$(GREEN)База данных сброшена$(NC)"

# Команды для API
api-test: ## Протестировать API
	@echo "$(GREEN)Тестирование API...$(NC)"
	@echo "$(YELLOW)Health check:$(NC)"
	curl -s http://localhost:8080/health | jq .
	@echo "$(YELLOW)Создание заявки:$(NC)"
	curl -s -X POST http://localhost:8080/api/v1/issue \
		-H "Content-Type: application/json" \
		-d '{"full_name": "Иван Иванов", "contact_info": "+7-999-123-45-67", "preferred_contact_method": "Телефон", "has_china_experience": true, "has_supplier_contacts": false, "product_description": "Электронные компоненты", "existing_product_links": "https://ozon.ru/product1", "expected_delivery_date": "2024-12-01"}' | jq .

# Команды для релиза
release: clean build ## Создать релиз
	@echo "$(GREEN)Релиз создан: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

# Команды для управления сервером
stop: ## Остановить сервер
	@echo "$(GREEN)Остановка сервера...$(NC)"
	@if pgrep -f "go run.*main.go" > /dev/null; then \
		pkill -f "go run.*main.go"; \
		echo "$(GREEN)Сервер остановлен$(NC)"; \
	elif pgrep -f "main" > /dev/null; then \
		pkill -f "main"; \
		echo "$(GREEN)Сервер остановлен$(NC)"; \
	else \
		echo "$(YELLOW)Сервер не запущен$(NC)"; \
	fi

.PHONY: help build run test clean deps fmt lint dev install-air docker-build docker-run db-migrate db-reset api-test release stop 
# Calc Example - Бэкенд на Go

Современный REST API бэкенд на языке Go с использованием Clean Architecture, GORM, Gin и SQLite.

## 🚀 Возможности

- **Clean Architecture** - Чистая архитектура с разделением на слои
- **REST API** - Полноценный REST API с валидацией
- **База данных** - SQLite с автоматической миграцией
- **Логирование** - Структурированное логирование с logrus
- **Конфигурация** - Гибкая конфигурация через переменные окружения
- **Graceful Shutdown** - Корректное завершение работы сервера
- **CORS** - Поддержка CORS для фронтенда
- **Health Check** - Проверка состояния сервера

## 📁 Структура проекта

```
calc_example/
├── cmd/
│   └── server/
│       └── main.go          # Точка входа приложения
├── internal/
│   ├── app/
│   │   └── app.go          # Основное приложение
│   ├── config/
│   │   └── config.go       # Конфигурация
│   ├── handler/
│   │   └── handler.go      # HTTP хендлеры
│   ├── model/
│   │   ├── user.go         # Модель пользователя
│   │   └── calculation.go  # Модель расчета
│   ├── repository/
│   │   └── repository.go   # Слой доступа к данным
│   └── service/
│       └── service.go      # Бизнес-логика
├── pkg/
│   ├── database/
│   │   └── database.go     # Работа с БД
│   └── logger/
│       └── logger.go       # Логирование
├── go.mod                  # Зависимости
├── .env.example           # Пример конфигурации
└── README.md              # Документация
```

## 🛠 Установка и запуск

### Требования

- Go 1.22 или выше
- Git
- Docker (опционально)

### Установка

#### Локальная установка

1. Клонируйте репозиторий:

```bash
git clone <repository-url>
cd calc_example
```

2. Установите зависимости:

```bash
go mod tidy
```

3. Создайте файл конфигурации:

```bash
cp env.example .env
```

4. Запустите сервер:

```bash
go run cmd/server/main.go
```

#### Использование Makefile

```bash
# Установка зависимостей
make deps

# Запуск сервера
make run

# Сборка приложения
make build

# Запуск тестов
make test

# Показать все команды
make help
```

#### Использование Docker

```bash
# Сборка и запуск с Docker Compose
docker-compose up --build

# Или только сборка образа
docker build -t calc_example .

# Запуск контейнера
docker run -p 8080:8080 calc_example
```

Сервер будет доступен по адресу: `http://localhost:8080`

## 📋 API Endpoints

### Пользователи

- `POST /api/v1/users/` - Создать пользователя
- `GET /api/v1/users/` - Получить всех пользователей
- `GET /api/v1/users/:id` - Получить пользователя по ID
- `PUT /api/v1/users/:id` - Обновить пользователя
- `DELETE /api/v1/users/:id` - Удалить пользователя

### Расчеты

- `POST /api/v1/calculations/` - Создать расчет
- `GET /api/v1/calculations/` - Получить все расчеты
- `GET /api/v1/calculations/:id` - Получить расчет по ID
- `DELETE /api/v1/calculations/:id` - Удалить расчет
- `GET /api/v1/users/:user_id/calculations` - Получить расчеты пользователя

### Система

- `GET /health` - Проверка состояния сервера

## 📝 Примеры запросов

### Создание пользователя

```bash
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Создание расчета

```bash
curl -X POST http://localhost:8080/api/v1/calculations/ \
  -H "Content-Type: application/json" \
  -d '{
    "operation": "add",
    "operand1": 10,
    "operand2": 5
  }'
```

### Получение всех пользователей

```bash
curl http://localhost:8080/api/v1/users/
```

## ⚙️ Конфигурация

Создайте файл `.env` в корне проекта:

```env
# Сервер
SERVER_PORT=8080
SERVER_HOST=localhost

# База данных
DB_DRIVER=sqlite
DB_NAME=calc_example

# Логирование
LOG_LEVEL=info
```

## 🧪 Тестирование

Запуск тестов:

```bash
go test ./...
```

## 🔧 Разработка

### Добавление новых моделей

1. Создайте модель в `internal/model/`
2. Добавьте методы в `internal/repository/`
3. Добавьте бизнес-логику в `internal/service/`
4. Создайте хендлеры в `internal/handler/`

### Добавление middleware

Добавьте middleware в функцию `setupMiddleware` в `internal/app/app.go`

## 📦 Сборка

Сборка для продакшена:

```bash
go build -o bin/server cmd/server/main.go
```

## 🚀 Деплой

1. Соберите приложение
2. Скопируйте бинарный файл на сервер
3. Настройте переменные окружения
4. Запустите приложение

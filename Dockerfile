# Многоэтапная сборка для оптимизации размера образа
# Используем Debian-based образ для лучшей совместимости с SQLite3
FROM golang:1.22-bullseye AS builder

# Установка необходимых пакетов
RUN apt-get update && apt-get install -y \
    git \
    ca-certificates \
    tzdata \
    gcc \
    g++ \
    libc6-dev \
    libsqlite3-dev \
    && rm -rf /var/lib/apt/lists/*

# Установка рабочей директории
WORKDIR /app

# Копирование файлов зависимостей
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Финальный образ
FROM debian:bullseye-slim

# Установка необходимых пакетов
RUN apt-get update && apt-get install -y \
    ca-certificates \
    tzdata \
    && rm -rf /var/lib/apt/lists/*

# Создание пользователя и группы для безопасности
RUN groupadd -r -g 1001 appgroup && \
    useradd -r -u 1001 -g appgroup -s /sbin/nologin appuser

# Установка рабочей директории
WORKDIR /app

# Копирование бинарного файла из builder
COPY --from=builder /app/main .

# Копирование файла конфигурации
COPY env.example .env

# Изменение владельца файлов
RUN chown -R appuser:appgroup /app

# Переключение на непривилегированного пользователя
USER appuser

# Открытие порта
EXPOSE 8080

# Проверка здоровья
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Запуск приложения
CMD ["./main"] 
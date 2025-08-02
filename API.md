# API Документация

## Базовый URL
```
http://localhost:8080
```

## Аутентификация
В текущей версии аутентификация не реализована. Все запросы выполняются без токенов.

## Endpoints

### Health Check

#### GET /health
Проверка состояния сервера.

**Ответ:**
```json
{
  "status": "ok",
  "message": "Сервер работает"
}
```

### Пользователи

#### POST /api/v1/users/
Создание нового пользователя.

**Тело запроса:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Ответ:**
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john@example.com",
  "created_at": "2025-08-02T18:22:56.318403+03:00",
  "updated_at": "2025-08-02T18:22:56.318403+03:00"
}
```

#### GET /api/v1/users/
Получение списка всех пользователей.

**Ответ:**
```json
[
  {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "created_at": "2025-08-02T18:22:56.318403+03:00",
    "updated_at": "2025-08-02T18:22:56.318403+03:00"
  }
]
```

#### GET /api/v1/users/:id
Получение пользователя по ID.

**Параметры:**
- `id` (path) - ID пользователя

**Ответ:**
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john@example.com",
  "created_at": "2025-08-02T18:22:56.318403+03:00",
  "updated_at": "2025-08-02T18:22:56.318403+03:00"
}
```

#### PUT /api/v1/users/:id
Обновление пользователя.

**Параметры:**
- `id` (path) - ID пользователя

**Тело запроса:**
```json
{
  "username": "new_username",
  "email": "new@example.com",
  "password": "newpassword123"
}
```

**Ответ:**
```json
{
  "id": 1,
  "username": "new_username",
  "email": "new@example.com",
  "created_at": "2025-08-02T18:22:56.318403+03:00",
  "updated_at": "2025-08-02T18:23:00.000000+03:00"
}
```

#### DELETE /api/v1/users/:id
Удаление пользователя.

**Параметры:**
- `id` (path) - ID пользователя

**Ответ:**
```json
{
  "message": "Пользователь успешно удален"
}
```

### Расчеты

#### POST /api/v1/calculations/
Создание нового расчета.

**Тело запроса:**
```json
{
  "operation": "add",
  "operand1": 10,
  "operand2": 5
}
```

**Поддерживаемые операции:**
- `add` - сложение
- `subtract` - вычитание
- `multiply` - умножение
- `divide` - деление

**Ответ:**
```json
{
  "id": 1,
  "user_id": 1,
  "operation": "add",
  "operand1": 10,
  "operand2": 5,
  "result": 15,
  "created_at": "2025-08-02T18:23:02.442857+03:00",
  "updated_at": "2025-08-02T18:23:02.442857+03:00"
}
```

#### GET /api/v1/calculations/
Получение списка всех расчетов.

**Ответ:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "operation": "add",
    "operand1": 10,
    "operand2": 5,
    "result": 15,
    "created_at": "2025-08-02T18:23:02.442857+03:00",
    "updated_at": "2025-08-02T18:23:02.442857+03:00"
  }
]
```

#### GET /api/v1/calculations/:id
Получение расчета по ID.

**Параметры:**
- `id` (path) - ID расчета

**Ответ:**
```json
{
  "id": 1,
  "user_id": 1,
  "operation": "add",
  "operand1": 10,
  "operand2": 5,
  "result": 15,
  "created_at": "2025-08-02T18:23:02.442857+03:00",
  "updated_at": "2025-08-02T18:23:02.442857+03:00"
}
```

#### DELETE /api/v1/calculations/:id
Удаление расчета.

**Параметры:**
- `id` (path) - ID расчета

**Ответ:**
```json
{
  "message": "Расчет успешно удален"
}
```

### Расчеты пользователя

#### GET /user-calculations/:user_id
Получение всех расчетов конкретного пользователя.

**Параметры:**
- `user_id` (path) - ID пользователя

**Ответ:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "operation": "add",
    "operand1": 10,
    "operand2": 5,
    "result": 15,
    "created_at": "2025-08-02T18:23:02.442857+03:00",
    "updated_at": "2025-08-02T18:23:02.442857+03:00"
  }
]
```

## Коды ошибок

### 400 Bad Request
Некорректные данные запроса или валидация не прошла.

**Пример:**
```json
{
  "error": "Неверные данные запроса"
}
```

### 404 Not Found
Ресурс не найден.

**Пример:**
```json
{
  "error": "Пользователь не найден"
}
```

### 500 Internal Server Error
Внутренняя ошибка сервера.

**Пример:**
```json
{
  "error": "Внутренняя ошибка сервера"
}
```

## Примеры использования

### Создание пользователя и расчетов

```bash
# 1. Создание пользователя
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123"
  }'

# 2. Создание расчета
curl -X POST http://localhost:8080/api/v1/calculations/ \
  -H "Content-Type: application/json" \
  -d '{
    "operation": "multiply",
    "operand1": 7,
    "operand2": 8
  }'

# 3. Получение всех расчетов пользователя
curl http://localhost:8080/user-calculations/1
```

### Получение всех данных

```bash
# Получение всех пользователей
curl http://localhost:8080/api/v1/users/

# Получение всех расчетов
curl http://localhost:8080/api/v1/calculations/

# Проверка состояния сервера
curl http://localhost:8080/health
``` 
# Todo App

REST API сервис для управления списками задач с авторизацией пользователей.

## Стек технологий
- Go
- Gin (HTTP framework)
- PostgreSQL
- JWT для авторизации
- Viper и dotenv для конфигурации
- sqlx для работы с базой данных
- Docker

## Функционал
- Регистрация и авторизация пользователей
- CRUD для списков задач (`TodoList`)
- CRUD для элементов внутри списков (`TodoItem`)
- JWT-аутентификация для защиты API
- Взаимодействие с PostgreSQL через транзакции
- Структурированная архитектура: handler → service → repository
- Разворачивается через Docker

## Установка

1. Клонировать репозиторий:
```bash
git clone https://github.com/Alekcey5977/todo-app
cd todo-app
```

Запуск приложения:
```bash
go run cmd/main.go
```

## Эндпоинты

### Auth
POST /auth/sign-up — регистрация пользователя

POST /auth/sign-in — авторизация, возвращает JWT токен

### Todo Lists (требуется JWT)
POST /api/lists/ — создать список

GET /api/lists/ — получить все списки пользователя

GET /api/lists/:id — получить список по ID

PUT /api/lists/:id — обновить список

DELETE /api/lists/:id — удалить список

### Todo Items
POST /api/lists/:id/items/ — создать элемент

GET /api/lists/:id/items/ — получить все элементы списка

GET /api/lists/:id/items/:item_id — получить элемент по ID

PUT /api/lists/:id/items/:item_id — обновить элемент

DELETE /api/lists/:id/items/:item_id — удалить элемент

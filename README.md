# Техническое задание: сервис управления задачами

## *tm - task manager

## Структура проекта

```
.
├── cmd/
│   └── main.go         # Точка входа
├── internal/
├── go.mod              # Зависимости Go-модуля
└── .env                # Переменные окружения
```

## Установка зависимостей проекта

```bash
go mod tidy
npm install
```

### Генерация с open api
```bash
go generate
```

### Поднять сервис в docker контейнере
```bash
docker-compose up -d
```

## Запуск проекта

```bash
go run cmd/main.go
```

## [Swagger🔍](http://localhost:8080/docs/index.html)



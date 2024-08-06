# Проект Agent REST API

## Используемые технологии и компоненты

- **Язык программирования**: Go (Golang)
- **Фреймворк для веб-сервера**: Gin
- **База данных**: PostgreSQL
- **Библиотеки и зависимости**:
  - `github.com/gin-gonic/gin`
  - `github.com/lib/pq`
  - `github.com/joho/godotenv`
  - `github.com/caarlos0/env/v9`

- **Схема SQL:** queries/schema.sql

## Шаги для запуска проекта

### 1. Установка Go

Убедитесь, что у вас установлена последняя версия Go. Вы можете скачать её с [официального сайта](https://golang.org/dl/).

### 2. Клонирование репозитория

Клонируйте репозиторий на ваш локальный компьютер:

- ``git clone https://github.com/yourusername/agent-rest-api.git``
- ``cd agent-rest-api``

### 3. Установка зависимостей

Установите все необходимые зависимости:

- ``go mod download``

### 4. Настройка базы данных
- Убедитесь, что PostgreSQL установлен и запущен.
- Создайте базу данных agent-db.
- Настройте параметры подключения в файле .env.local и config.json.
   #### .env.local
  - POSTGRES_USER=your_postgres_user
  - POSTGRES_PASSWORD=your_postgres_password
  - DOMAIN=http://localhost:3000 (пример)
   #### config.json
  {
    "postgres": {
        "db_host": "localhost",
        "db_port": "5432",
        "db_name": "agent-db",
        "ssl_mode": "disable"
    },
    "server": {
        "port": "8080"
    }
}
### 5. Запуск проекта

Запустите проект с помощью следующей команды:

- `go run cmd/main.go`

### 6. Проверка работы API

Код для тела запроса Postman:

{
    "name": "Agent Smith",
    "status": "Active"
}

## Структура проекта

```plaintext
agent-rest-api/
├── cmd/
│   └── main.go
├── configs/
│   └── config.go
├── models/
│   └── env.go
│   └── agent.go
│   └── task.go
├── controllers/
│   └── agent_controller.go
├── repositories/
│   └── database.go
│   └── agent_repository.go
├── services/
│   └── agent_service.go
├── handlers/
│   └── router.go
├── go.mod
├── go.sum
└── .env.local

# Messageio

Messageio — это проект, предназначенный для обработки и хранения сообщений с использованием Kafka и PostgreSQL. Проект включает в себя HTTP API для отправки и получения сообщений, а также обработку сообщений из Kafka.


## Требования

- Go 1.20 или выше
- Docker
- Docker Compose

## Установка

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/Chigvero/Messageio.git
    cd Messageio
    ```


2. Соберите и запустите проект с помощью Docker Compose:

    ```sh
    docker-compose up --build
    ```

## Конфигурация

Конфигурация проекта находится в файле `configs/config.yml`. Пример конфигурации:

```yaml
db:
  host: "db"
  port: "5433"
  username: "postgres"
  password: "1111"
  dbname: "Intern"
  ssl_mode: "disable"
```
## Cтруктура базы данных
```
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    processed BOOLEAN NOT NULL DEFAULT false
);
```

## API

`api/v1/message/:id` GET получение сообщения по id


`api/v1/message/` POST отправка сообщения

`api/v1/stats` GET получение статистики по обработанным сообзениям
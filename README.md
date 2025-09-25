# Steam Inventory

Демо-сервис: авторизация через **Steam OpenID**, хранение сессий в **Redis**, база пользователей в **PostgreSQL**.

### Возможности на данный момент:
- Логин через Steam OpenID
- Сессии (cookie) с хранением в Redis
- Создание/обновление пользователя в БД (SteamID, ник, аватар)

Конфиг можно задать через YAML, .env, env-переменные и флаги.
Описание полей в [/internal/config/types.go](/internal/config/types.go)

### Запуск

```sh
# поднять Postgres и Redis
docker compose up -d

# применить миграции
make migrate    # или вручную: goose -dir ./migrations postgres "$POSTGRES_DSN" up

# запустить бэкенд
make run
```

Стек: Go, Redis, PostgreSQL, Docker Compose

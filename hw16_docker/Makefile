# Проверка наличия файла .env и его включение
ifneq ($(wildcard .env),)
	include .env
endif

PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)\bin
PROJECT_TMP = $(PROJECT_DIR)\tmp

# Установка зависимостей
install-deps:
	GOBIN=$(PROJECT_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(PROJECT_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Статус миграции
migration-status:
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

# Добавление миграции
migration-add:
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} create $(name) sql

# Применение миграций
migration-up:
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

# Откат миграций
migration-down:
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

# Генерация SQL
gen-sql:
	rm -f internal/repository/product/*.go
	$(PROJECT_BIN)/sqlc generate
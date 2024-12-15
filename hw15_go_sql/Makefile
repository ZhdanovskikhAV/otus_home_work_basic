# Проверка наличия файла .env и его включение
ifneq ($(wildcard .env),)
	include .env
endif

PROJECT_DIR = E:\work\GO\otus_home_work_basic\hw15_go_sql
PROJECT_BIN = $(PROJECT_DIR)\bin
PROJECT_TMP = $(PROJECT_DIR)\tmp

# Установка зависимостей
install-deps:
	$env:GOBIN="$(PROJECT_BIN)"; go install github.com/pressly/goose/v3/cmd/goose@latest
	$env:GOBIN="$(PROJECT_BIN)"; go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Статус миграции
migration-status:
	powershell -Command "& { & '$(PROJECT_BIN)\goose' -dir '${MIGRATION_DIR}' postgres '${MIGRATION_DSN}' status -v }"

# Добавление миграции
migration-add:
	powershell -Command "& { & '$(PROJECT_BIN)\goose' -dir '${MIGRATION_DIR}' create '$(name)' sql }"

# Применение миграций
migration-up:
	powershell -Command "& { & '$(PROJECT_BIN)\goose' -dir '${MIGRATION_DIR}' postgres '${MIGRATION_DSN}' up -v }"

# Откат миграций
migration-down:
	powershell -Command "& { & '$(PROJECT_BIN)\goose' -dir '${MIGRATION_DIR}' postgres '${MIGRATION_DSN}' down -v }"

# Генерация SQL
gen-sql:
	powershell -Command "& { Remove-Item -Path 'internal\repository\product\*.go' -Force }"
	powershell -Command "& { & '$(PROJECT_BIN)\sqlc' generate }"

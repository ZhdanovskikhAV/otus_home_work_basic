# Используем образ Golang для сборки приложения
FROM golang:1.23.1 as builder

# Устанавливаем рабочую директорию
WORKDIR /build/hw16_docker/

# Копируем .dockerignore и файлы зависимостей
COPY dockerignore ./

COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все остальные файлы в рабочую директорию
COPY . .

# Сборка приложения с отключенным CGO
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server_api ./main.go

# Используем легковесный образ Alpine для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию для финального контейнера
WORKDIR /root/

# Копируем скомпилированное приложение из этапа сборки
COPY --from=builder /build/hw16_docker/server_api ./

# Устанавливаем переменную окружения
ENV CONFIG_ENV=".env"

# Открываем необходимые порты
EXPOSE 8080
EXPOSE 5432

# Указываем команду для запуска приложения
CMD ["./server_api"]
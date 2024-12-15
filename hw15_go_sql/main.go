package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw15_go_sql/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получаем значения переменных окружения
	dsn := os.Getenv("DB_DSN")
	ip := os.Getenv("APP_HOST")
	port := os.Getenv("APP_HTTP_PORT")
	fmt.Printf("DB_DSN: %s\n", dsn)
	fmt.Printf("APP_HOST: %s\n", ip)
	fmt.Printf("APP_HTTP_PORT: %s\n", port)
	if len(dsn) == 0 || len(ip) == 0 || len(port) == 0 {
		fmt.Println("DB_DSN or APP_HOST or APP_HTTP_PORT not set")
		return
	}

	server.Server(&ip, &port, &dsn)
}

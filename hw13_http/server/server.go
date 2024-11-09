package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Printf("Получен GET запрос: %s\n", r.URL.Path)
		w.Write([]byte("Hello from the server!"))
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Printf("Получен POST запрос: %s\n", r.URL.Path)
		fmt.Printf("Данные POST запроса: %s\n", string(body))
		w.Write([]byte("Data received!"))
	}
}

func main() {
	// Определяем флаги для IP и порта
	ip := flag.String("ip", "localhost", "IP адрес для сервера")
	port := flag.String("port", "8080", "Порт для сервера")

	// Парсим флаги
	flag.Parse()

	// Формируем адрес сервера
	address := fmt.Sprintf("%s:%s", *ip, *port)

	http.HandleFunc("/", handler)

	// Создаем новый HTTP сервер с таймаутами
	server := &http.Server{
		Addr:         address,
		Handler:      nil,               // Используем nil для стандартного маршрутизатора
		ReadTimeout:  10 * time.Second,  // Таймаут для чтения
		WriteTimeout: 10 * time.Second,  // Таймаут для записи
		IdleTimeout:  120 * time.Second, // Таймаут для бездействия
	}

	fmt.Printf("Сервер запущен на %s...\n", address)
	log.Fatal(server.ListenAndServe()) // Запускаем сервер
}

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// isValidURL проверяет, является ли URL допустимым HTTP или HTTPS.
func isValidURL(rawURL string) (*url.URL, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	// Проверка схемы URL
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, fmt.Errorf("недопустимый протокол: %s", parsedURL.Scheme)
	}
	return parsedURL, nil
}

// sendRequest отправляет HTTP запрос на указанный URL с данными и методом.
func sendRequest(method, requestURL string, data string) {
	parsedURL, err := isValidURL(requestURL)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Создаем контекст с тайм-аутом.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, parsedURL.String(), bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Printf("Ответ на %s запрос: %s\n", method, string(body))
}

// main функция для запуска клиента.
func main() {
	urlFlag := flag.String("url", "", "URL сервера")
	postData := flag.String("post", "", "Данные для отправки с POST запросом")
	putData := flag.String("put", "", "Данные для отправки с PUT запросом")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Ошибка: необходимо указать URL")
		return
	}

	switch {
	case *putData != "":
		sendRequest(http.MethodPut, *urlFlag, *putData)
	case *postData != "":
		sendRequest(http.MethodPost, *urlFlag, *postData)
	default:
		sendRequest(http.MethodGet, *urlFlag, "")
	}
}

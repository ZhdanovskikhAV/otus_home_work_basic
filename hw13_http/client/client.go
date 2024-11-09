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

// sendGetRequest отправляет GET запрос на указанный URL.
func sendGetRequest(requestURL string) {
	parsedURL, err := isValidURL(requestURL)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Создаем контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ на GET запрос:", string(body))
}

// sendPostRequest отправляет POST запрос на указанный URL с данными.
func sendPostRequest(requestURL string, data string) {
	parsedURL, err := isValidURL(requestURL)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Создаем контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, parsedURL.String(), bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке POST запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ на POST запрос:", string(body))
}

// main функция для запуска клиента.
func main() {
	urlFlag := flag.String("url", "", "URL сервера")
	postData := flag.String("post", "", "Данные для отправки с POST запросом")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Ошибка: необходимо указать URL")
		return
	}

	if *postData != "" {
		sendPostRequest(*urlFlag, *postData)
	} else {
		sendGetRequest(*urlFlag)
	}
}

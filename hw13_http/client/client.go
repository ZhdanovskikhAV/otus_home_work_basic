package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func isValidURL(rawURL string) bool {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	// Проверка схемы URL
	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}

func sendGetRequest(url string) {
	if !isValidURL(url) {
		fmt.Println("Ошибка: недопустимый URL:", url)
		return
	}

	resp, err := http.Get(url)
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

func sendPostRequest(url string, data string) {
	if !isValidURL(url) {
		fmt.Println("Ошибка: недопустимый URL:", url)
		return
	}

	resp, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(data)))
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

func main() {
	url := flag.String("url", "", "URL сервера")
	postData := flag.String("post", "", "Данные для отправки с POST запросом")
	flag.Parse()

	if *url == "" {
		fmt.Println("Ошибка: необходимо указать URL")
		return
	}

	if *postData != "" {
		sendPostRequest(*url, *postData)
	} else {
		sendGetRequest(*url)
	}
}

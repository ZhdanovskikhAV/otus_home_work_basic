package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSendGetRequest(t *testing.T) {
	// Создаем тестовый сервер.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello from test server!"))
	}))
	defer ts.Close()

	// Создаем контекст с таймаутом.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Выполняем GET запрос с использованием контекста.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatalf("Ошибка при создании GET запроса: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка при отправке GET запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Неверный статус-код: ожидался %v, получен %v", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Ошибка при чтении тела ответа: %v", err)
	}
	expected := "Hello from test server!"
	if string(body) != expected {
		t.Errorf("Неверное тело ответа: ожидалось %v, получено %v", expected, string(body))
	}
}

func TestSendPostRequest(t *testing.T) {
	// Создаем тестовый сервер.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		if string(body) == "Hello Server!" {
			w.Write([]byte("Data received!"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer ts.Close()

	// Создаем контекст с таймаутом.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Выполняем POST запрос с использованием контекста.
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ts.URL, bytes.NewBuffer([]byte("Hello Server!")))
	if err != nil {
		t.Fatalf("Ошибка при создании POST запроса: %v", err)
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка при отправке POST запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Неверный статус-код: ожидался %v, получен %v", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Ошибка при чтении тела ответа: %v", err)
	}
	expected := "Data received!"
	if string(body) != expected {
		t.Errorf("Неверное тело ответа: ожидалось %v, получено %v", expected, string(body))
	}
}

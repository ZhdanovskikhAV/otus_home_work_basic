package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandlerGet(t *testing.T) {
	// Создаем контекст с таймаутом.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный статус-код: ожидался %v, получен %v", http.StatusOK, status)
	}

	expected := "Hello from the server!"
	if rr.Body.String() != expected {
		t.Errorf("Неверное тело ответа: ожидалось %v, получено %v", expected, rr.Body.String())
	}
}

func TestHandlerPost(t *testing.T) {
	// Создаем контекст с таймаутом.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := []byte("Hello Server!")
	req, err := http.NewRequestWithContext(ctx, "POST", "/test", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный статус-код: ожидался %v, получен %v", http.StatusOK, status)
	}

	expected := "Data received!"
	if rr.Body.String() != expected {
		t.Errorf("Неверное тело ответа: ожидалось %v, получено %v", expected, rr.Body.String())
	}
}

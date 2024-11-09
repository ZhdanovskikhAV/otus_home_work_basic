package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
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
	payload := []byte("Hello Server!")
	req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(payload))
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

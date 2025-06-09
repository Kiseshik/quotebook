package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestAddQuoteHandler(t *testing.T) {
	storage := NewStorage()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddQuoteHandler(w, r, storage)
	})

	// Тест успешного добавления
	reqBody := `{"author": "Test Author", "quote": "Test Quote"}`
	req := httptest.NewRequest("POST", "/quotes", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	// Тест с неверным телом запроса
	req = httptest.NewRequest("POST", "/quotes", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetAllQuotesHandler(t *testing.T) {
	storage := NewStorage()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetAllQuotesHandler(w, r, storage)
	})

	// Добавляем тестовые данные
	storage.AddQuote("Author1", "Quote1")
	storage.AddQuote("Author2", "Quote2")

	req := httptest.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	if !strings.Contains(w.Body.String(), "Quote1") || !strings.Contains(w.Body.String(), "Quote2") {
		t.Error("Expected response to contain test quotes")
	}
}

func TestDeleteQuoteHandler(t *testing.T) {
	storage := NewStorage()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteQuoteHandler(w, r, storage)
	})

	// Добавляем тестовые данные
	quote := storage.AddQuote("Author", "Quote")

	req := httptest.NewRequest("DELETE", "/quotes/"+strconv.Itoa(quote.ID), nil)
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status %d, got %d", http.StatusNoContent, w.Code)
	}

	// Попытка удалить несуществующую цитату
	req = httptest.NewRequest("DELETE", "/quotes/999", nil)
	w = httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

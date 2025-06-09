package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func AddQuoteHandler(w http.ResponseWriter, r *http.Request, storage *Storage) {
	var req RequestQuote
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Author == "" || req.Text == "" {
		http.Error(w, "Author and quote text are required", http.StatusBadRequest)
		return
	}

	quote := storage.AddQuote(req.Author, req.Text)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func GetAllQuotesHandler(w http.ResponseWriter, r *http.Request, storage *Storage) {
	quotes := storage.GetAllQuotes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func GetQuotesByAuthorHandler(w http.ResponseWriter, r *http.Request, storage *Storage) {
	author := r.URL.Query().Get("author")
	if author == "" {
		http.Error(w, "Author parameter is required", http.StatusBadRequest)
		return
	}

	quotes := storage.GetQuotesByAuthor(author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func GetRandomQuoteHandler(w http.ResponseWriter, r *http.Request, storage *Storage) {
	quote, ok := storage.GetRandomQuote()
	if !ok {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func DeleteQuoteHandler(w http.ResponseWriter, r *http.Request, storage *Storage) {
	path := strings.TrimPrefix(r.URL.Path, "/quotes/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid quote ID", http.StatusBadRequest)
		return
	}

	if !storage.DeleteQuote(id) {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package main

import (
	"log"
	"net/http"
)

func main() {
	storage := NewStorage()

	http.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if author := r.URL.Query().Get("author"); author != "" {
				GetQuotesByAuthorHandler(w, r, storage)
				return
			}
			GetAllQuotesHandler(w, r, storage)
		case http.MethodPost:
			AddQuoteHandler(w, r, storage)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/quotes/random", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetRandomQuoteHandler(w, r, storage)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/quotes/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			DeleteQuoteHandler(w, r, storage)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

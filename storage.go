package main

import (
	"math/rand"
	"slices"
	"sync"
	"time"
)

type Storage struct {
	quotes []Quote
	mu     sync.RWMutex
	nextID int
}

func NewStorage() *Storage {
	return &Storage{
		quotes: make([]Quote, 0),
		nextID: 1,
	}
}

func (s *Storage) AddQuote(author, text string) Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	quote := Quote{
		ID:        s.nextID,
		Author:    author,
		Text:      text,
		CreatedAt: time.Now(),
	}

	s.quotes = append(s.quotes, quote)
	s.nextID++
	return quote
}

func (s *Storage) GetAllQuotes() []Quote {
	s.mu.RLock()
	defer s.mu.RUnlock()

	quotes := make([]Quote, len(s.quotes))
	copy(quotes, s.quotes)
	return quotes
}

func (s *Storage) GetQuotesByAuthor(author string) []Quote {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []Quote
	for _, quote := range s.quotes {
		if quote.Author == author {
			result = append(result, quote)
		}
	}
	return result
}

func (s *Storage) GetRandomQuote() (Quote, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.quotes) == 0 {
		return Quote{}, false
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return s.quotes[r.Intn(len(s.quotes))], true
}

func (s *Storage) DeleteQuote(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, quote := range s.quotes {
		if quote.ID == id {
			s.quotes = slices.Delete(s.quotes, i, i+1)
			return true
		}
	}
	return false
}

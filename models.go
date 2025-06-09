package main

import "time"

type Quote struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Text      string    `json:"quote"`
	CreatedAt time.Time `json:"created_at"`
}

type RequestQuote struct {
	Author string `json:"author"`
	Text   string `json:"quote"`
}

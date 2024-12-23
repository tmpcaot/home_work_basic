package main

import (
	"encoding/json"
)

// JSONBook представляет собой структуру книги, предназначенную для сериализации и десериализации в JSON.
type JSONBook struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

// MarshalJSON переопределяет стандартную функцию маршалинга для JSONBook.
func (b JSONBook) MarshalJSON() ([]byte, error) {
	type Alias JSONBook
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(b),
	})
}

// UnmarshalJSON переопределяет стандартную функцию демаршалинга для JSONBook.
func (b *JSONBook) UnmarshalJSON(data []byte) error {
	type Alias JSONBook
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}

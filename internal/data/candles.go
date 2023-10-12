package data

import (
	"encoding/json"
	"time"
)

type Candle struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price,omitempty"`
}

func (c Candle) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price,omitempty"`
	}{
		// Set the values for the anonymous struct.
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Price:       c.Price,
	}
	// Encode the anonymous struct to JSON, and return it.
	return json.Marshal(aux)
}

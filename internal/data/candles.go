package data

import (
	"time"
)

type Candle struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price,omitempty"`
}

package data

import (
	"time"
)

type Candle struct {
	ID          int64 // Unique integer ID for the candle
	CreatedAt   time.Time
	Name        string
	Description string
	Price       float64
}

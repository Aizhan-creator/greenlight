package data

import (
	"encoding/json"
	"greenlight.alexedwards.net/internal/validator"
	"time"
)

type Candle struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price,omitempty"`
	Runtime     Runtime   `json:"runtime,omitempty"`
}

func ValidateCandle(v *validator.Validator, candle *Candle) {
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(input.Description != "", "description", "must be provided")
	v.Check(len(input.Description) <= 500, "description", "must not be more than 500 bytes long")

	v.Check(input.Runtime != 0, "runtime", "must be provided")
	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(input.Price != 0, "price", "must be provided")
	v.Check(input.Price > 0, "price", "must be a positive integer")
}

func (c Candle) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price,omitempty"`
	}{

		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Price:       c.Price,
	}
	return json.Marshal(aux)
}

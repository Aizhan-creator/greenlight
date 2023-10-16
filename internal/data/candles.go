package data

import (
	"encoding/json"
	"greenlight.alexedwards.net/internal/validator"
	//"github.com/Aizhan-creator/greenlight/internal/validator"
	"time"
)

type Candle struct {
	ID          int64     //`json:"id"`
	CreatedAt   time.Time //`json:"-"`
	Name        string    //`json:"name"`
	Description string    //`json:"description"`
	Price       float64   //`json:"price,omitempty"`
	Runtime     Runtime   //`json:"runtime,omitempty"`
}

func ValidateCandle(v *validator.Validator, candle *Candle) {

	v.Check(candle.Name != "", "name", "must be provided")
	v.Check(len(candle.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(candle.Description != "", "description", "must be provided")
	v.Check(len(candle.Description) <= 500, "description", "must not be more than 500 bytes long")

	v.Check(candle.Runtime != 0, "runtime", "must be provided")
	v.Check(candle.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(candle.Price != 0, "price", "must be provided")
	v.Check(candle.Price > 0, "price", "must be a positive integer")
}
func (c Candle) MarshalJSON() ([]byte, error) {
	var runtime string

	aux := struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Runtime     string  `json:"runtime,omitempty"`
		Price       float64 `json:"price,omitempty"`
	}{

		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Runtime:     runtime,
		Price:       c.Price,
	}
	return json.Marshal(aux)
}

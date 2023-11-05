package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/lib/pq"
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

type MockCandleModel struct{}

func (c MockCandleModel) Insert(candle *Candle) error {
	return nil
}

func (c MockCandleModel) Get(id int64) (*Candle, error) {
	return nil, nil
}

func (c MockCandleModel) Update(candle *Candle) error {
	return nil
}

func (c MockCandleModel) Delete(id int64) error {
	return nil
}

type CandleModel struct {
	DB *sql.DB
}

func (c CandleModel) Insert(candle *Candle) error {
	query := `INSERT INTO candles (name, description, runtime, price)
				VALUES ($1, $2, $3, $4)
				RETURNING id, created_at`
	args := []interface{}{candle.Name, candle.Description, candle.Runtime, pq.Array(candle.Price)}
	return c.DB.QueryRow(query, args...).Scan(&candle.ID, &candle.CreatedAt)

}
func (c CandleModel) Get(id int64) (*Candle, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT id, created_at, name, description, runtime, price
				FROM candles
				WHERE id = $1`

	var candle Candle

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, id).Scan(
		&candle.ID,
		&candle.CreatedAt,
		&candle.Name,
		&candle.Description,
		&candle.Runtime,
		&candle.Price,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &candle, nil
}

func (c CandleModel) Update(candle *Candle) error {
	query := `
		UPDATE candles
		SET name = $1, description = $2, runtime = $3, price = $4
		WHERE id = $5 
		RETURNING version`

	args := []interface{}{
		candle.Name,
		candle.Description,
		candle.Runtime,
		candle.ID,
		candle.Price,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&candle.Price)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (c CandleModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM candles
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := c.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

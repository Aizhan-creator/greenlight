package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Candles interface {
		Insert(movie *Candle) error
		Get(id int64) (*Candle, error)
		Update(movie *Candle) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Candles: CandleModel{DB: db},
	}
}
func NewMockModels() Models {
	return Models{
		Candles: MockCandleModel{},
	}
}

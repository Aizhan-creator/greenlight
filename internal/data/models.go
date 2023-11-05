package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies CandleModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: CandleModel{DB: db},
	}
}

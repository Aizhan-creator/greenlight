package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Candles     CandleModel
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Candles:     CandleModel{DB: db},
		Users:       UserModel{DB: db}, // Initialize a new UserModel instance.
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}

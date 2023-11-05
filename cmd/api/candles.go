package main

import (
	"fmt"
	"greenlight.alexedwards.net/internal/data"
	"greenlight.alexedwards.net/internal/validator"
	//"github.com/Aizhan-creator/greenlight/internal/data"
	//"github.com/Aizhan-creator/greenlight/internal/validator"
	"net/http"
	"time"
)

func (app *application) createCandleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Runtime     data.Runtime `json:"runtime"`
		Price       float64      `json:"price,omitempty"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	candle := &data.Candle{
		Name:        input.Name,
		Description: input.Description,
		Runtime:     input.Runtime,
		Price:       input.Price,
	}
	v := validator.New()

	if data.ValidateCandle(v, candle); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//fmt.Fprintf(w, "%+v\n", input)
	err = app.models.Candles.Insert(candle)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/candles/%d", candle.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"candle": candle}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) showCandleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	candle := data.Candle{
		ID:          id,
		CreatedAt:   time.Now(),
		Name:        "Pink",
		Description: "Candle with roses",
		Runtime:     102,
		Price:       4000.0,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"candle": candle}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

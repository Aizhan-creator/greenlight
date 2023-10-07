package main

import (
	"encoding/json"
	"fmt"
	"greenlight/internal/data"
	"net/http"
	"time" // New import
)

func (app *application) createCandleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new candle")
}
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price,omitempty"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
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
		Price:       4000.0,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"candle": candle}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

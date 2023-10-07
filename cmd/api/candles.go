package main

import (
	"fmt"
	"greenlight/internal/data"
	"net/http"
	"time" // New import
)

func (app *application) createCandleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new candle")
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
	err = app.writeJSON(w, http.StatusOK, candle, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

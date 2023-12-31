package main

import (
	"errors"
	"fmt"
	"greenlight.alexedwards.net/internal/data"
	"greenlight.alexedwards.net/internal/validator"
	"strconv"

	//"github.com/Aizhan-creator/greenlight/internal/data"
	//"github.com/Aizhan-creator/greenlight/internal/validator"
	"net/http"
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
	candle, err := app.models.Candles.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"candle": candle}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) updateCandleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	candle, err := app.models.Candles.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		app.logger.PrintInfo("error", nil)
		return
	}
	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(candle.Price), 32) != r.Header.Get("X-Expected-Version") {
			app.editConflictResponse(w, r)
			return
		}
	}
	var input struct {
		Name        *string       `json:"name"`
		Description *string       `json:"description"`
		Runtime     *data.Runtime `json:"runtime"`
		Price       *float64      `json:"price"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.logger.PrintInfo("wrong body params", nil)
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		candle.Name = *input.Name
	}
	// We also do the same for the other fields in the input struct.
	if input.Description != nil {
		candle.Description = *input.Description
	}
	if input.Runtime != nil {
		candle.Runtime = *input.Runtime
	}
	if input.Price != nil {
		candle.Price = *input.Price
	}

	v := validator.New()
	if data.ValidateCandle(v, candle); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Candles.Update(candle)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"candle": candle}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteCandleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Candles.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "candle successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listCandlesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string
		Description string
		data.Filters
	}
	v := validator.New()

	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Description = app.readString(qs, "description", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	// Read the sort query string value into the embedded struct.
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "name", "description", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	candles, metadata, err := app.models.Candles.GetAll(input.Name, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"candles": candles, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

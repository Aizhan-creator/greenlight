package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/candles", app.listCandleHandler)
	router.HandlerFunc(http.MethodPost, "/v1/candles", app.createCandleHandler)
	router.HandlerFunc(http.MethodGet, "/v1/candles/:id", app.showCandleHandler)
	router.HandlerFunc(http.MethodPut, "/v1/candles/:id", app.updateCandleHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/candles/:id", app.deleteCandleHandler)

	return router
}

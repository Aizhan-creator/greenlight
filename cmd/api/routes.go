package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/candles", app.listCandlesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/candles", app.createCandleHandler)
	router.HandlerFunc(http.MethodGet, "/v1/candles/:id", app.showCandleHandler)
	router.HandlerFunc(http.MethodPut, "/v1/candles/:id", app.updateCandleHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/candles/:id", app.deleteCandleHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.rateLimit(router))
}

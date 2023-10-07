package main

import (
	"encoding/json" // New import
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)
...
type envelope map[string]interface{}
// Change the data parameter to have the type envelope instead of interface{}.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

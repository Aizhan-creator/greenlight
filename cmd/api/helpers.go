package main

import (
	"encoding/json" // New import
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)
...
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	// Use the json.MarshalIndent() function so that whitespace is added to the encoded
	// JSON. Here we use no line prefix ("") and tab indents ("\t") for each element.
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
package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) Add(w http.ResponseWriter, r *http.Request) {
	// read json from the request body
	//var person functions.Person
	//_ = app.readJSON(w, r, person)
	// insert person into in-memory slice
	//app.People = functions.AddSinglePersonAndMatch(app.People, person)

	payload := jsonResponse{
		Error:   false,
		Message: "Add a person to the list",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

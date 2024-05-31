package main

import (
	"net/http"
)

func (app *application) healtcheck(w http.ResponseWriter, r *http.Request) {
	data := responseFormat{
		"status":      "running",
		"environment": app.config.env,
		"version":     version,
	}
	app.writeJSON(w, http.StatusOK, data, nil)
}

package main

import (
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message any) {
	err := app.writeJSON(w, statusCode, message, nil)
	if err != nil {
		app.logError(r, err)
	}

}

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

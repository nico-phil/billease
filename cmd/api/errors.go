package main

import (
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message any) {

	err := app.writeJSON(w, statusCode, responseFormat{"error": message}, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(status)
	}

}

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

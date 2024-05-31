package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healtcheck(w http.ResponseWriter, r *http.Request) {
	js, _ := json.MarshalIndent(map[string]string{"data": "Hello world"}, "", "\t")
	js = append(js, '\n')
	w.Write(js)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nico2220/billease/internal/data"
)

func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		From     int64          `json:"from"`
		To       int64          `json:"to"`
		Services []data.Service `json:"services"`
		Vat      int            `json:"vat"`
		Currency string         `json:"currency"`
		// SubTotal int            `json:"subTotal"`
		// Tax      int            `json:"tax"`
		// Total    int            `json:"total"`
	}

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&input)
	if err != nil {
		app.writeJSON(w, http.StatusInternalServerError, responseFormat{"error": err.Error()}, nil)
	}

	invoice := data.Invoice{
		From:     input.From,
		To:       input.To,
		Services: input.Services,
		Vat:      input.Vat,
	}

	invoice.CalculateSubTotal()
	invoice.CalculateTax()
	invoice.CalculateTotal()

	app.writeJSON(w, http.StatusOK, responseFormat{"data": "created"}, nil)

	fmt.Printf("%+v", invoice)

}

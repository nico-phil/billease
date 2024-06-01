package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nico2220/billease/internal/data"
	"github.com/Nico2220/billease/internal/pdf"
	"github.com/Nico2220/billease/internal/validator"
)

func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		From     int64          `json:"from"`
		To       int64          `json:"to"`
		Services []data.Service `json:"services"`
		Vat      int            `json:"vat"`
		Currency string         `json:"currency"`
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

	v := validator.New()

	if data.ValidateInvoice(v, &invoice); !v.Valid() {
		app.writeJSON(w, http.StatusInternalServerError, responseFormat{"error": "error append"}, nil)
	}

	// insert invoice into db

	// create pdf
	c1 := data.GetCompany(invoice.From)

	c2 := data.GetCompany(invoice.To)
	pdf.New(invoice, c1, c2)

	app.writeJSON(w, http.StatusOK, responseFormat{"data": "created"}, nil)

	fmt.Printf("%+v", invoice)

}

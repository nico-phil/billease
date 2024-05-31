package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nico2220/billease/internal/data"
	"github.com/Nico2220/billease/internal/pdf"
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

	// insert invoice into db

	// create pdf
	c1 := getCompany(invoice.From)

	c2 := getCompany(invoice.To)
	pdf.New(invoice, c1, c2)

	app.writeJSON(w, http.StatusOK, responseFormat{"data": "created"}, nil)

	fmt.Printf("%+v", invoice)

}

func getCompany(id int64) data.Company {
	if id == 1 {
		return data.Company{
			Name:         "PhiTech Nico",
			Contact:      "Nicolas",
			Adress:       "Germain street",
			Country:      "Estonia",
			SocityNumber: "123456",
			Code:         "585943",
			VatNumber:    "EE1234445",
			PhoneNumber:  "+79772820353",
			Email:        "nphilibert17@gmail.com",
		}
	}

	return data.Company{
		Name:         "The Good Seat",
		Contact:      "Alex",
		Adress:       "Rue des entrepreneur",
		Country:      "France",
		SocityNumber: "123456",
		Code:         "585943",
		VatNumber:    "EE1234445",
		PhoneNumber:  "+79772820353",
		Email:        "alex@gmail.com",
	}
}

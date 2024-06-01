package data

import (
	"github.com/Nico2220/billease/internal/validator"
	"time"
)

type Invoice struct {
	ID       int64
	From     int64
	To       int64
	Services []Service
	SubTotal float64
	Tax      float64
	Total    float64
	Currency string
	Vat      int
	CreateAt time.Time
	Link     string
}

type Service struct {
	ServiceType string  `json:"service_type"`
	Description string  `json:"description"`
	Rate        float64 `json:"rate"`
	Quantity    float64 `json:"quantity"`
	Amount      float64 `json:"amount"`
}

func ValidateInvoice(v *validator.Validator, invoice *Invoice) {
	v.Check(invoice.From > 0, "from", "from must be provided")
	v.Check(invoice.To > 0, "to", "to must be provided")
	v.Check(len(invoice.Services) > 0, "services", "must be containt at least 1 service")
	v.Check(invoice.Vat >= 0, "vat", "must be a posifive number")
	v.Check(invoice.Currency != "", "currency", "must be provided")

}

func (i *Invoice) CalculateSubTotal() {
	for _, v := range i.Services {
		i.SubTotal += v.Amount
	}
}

func (i *Invoice) CalculateTax() {
	i.Tax = (i.SubTotal * float64(i.Vat)) / 100
}

func (i *Invoice) CalculateTotal() {
	i.Total = i.Tax + i.SubTotal
}

type InvoiceModel struct {
	DB map[string]Invoice
}

func (m *InvoiceModel) Insert(invoice Invoice) (Invoice, error) {
	return invoice, nil

}

func GetCompany(id int64) Company {
	if id == 1 {
		return Company{
			Name:         "PhiTech",
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

	return Company{
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

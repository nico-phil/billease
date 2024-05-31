package data

import (
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
}

type Service struct {
	ServiceType string  `json:"service_type"`
	Description string  `json:"description"`
	Rate        float64 `json:"rate"`
	Quantity    float64 `json:"quantity"`
	Amount      float64 `json:"amount"`
}

// type InputData struct {
// 	Services []Service
// 	SubTotal int
// 	Tax      float
// 	Total    int
// 	Currency string
// 	Vat      int
// }

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
	DB []string
}

// func NewData(data InputData, from, to Company) *Invoice {
// 	return &Invoice{
// 		From:     from,
// 		To:       to,
// 		Services: data.Services,
// 		SubTotal: data.SubTotal,
// 		Tax:      data.Tax,
// 		Total:    data.Total,
// 		Currency: data.Currency,
// 		Vat: data.Vat,
// 	}
// }

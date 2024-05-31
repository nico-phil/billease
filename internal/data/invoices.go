package data

type Invoice struct {
	From     Company
	To       Company
	Services []Service
	SubTotal int
	Tax      int
	Total    int
	Currency string
	Vat      int
}

type Service struct {
	ServiceType string
	Description string
	Rate        int
	Quantity    int
	Amount      int
}

type InputData struct {
	Services []Service
	SubTotal int
	Tax      int
	Total    int
	Currency string
	Vat      int
}

func (i *Invoice) CalculateSubTotal() {
	for _, v := range i.Services {
		i.SubTotal += v.Amount
	}
}

func (i *Invoice) CalculateTax() {
	i.Tax = int((i.SubTotal * i.Vat) / 100)

}

func (i *Invoice) CalculateTotal() {
	i.Total = int(i.Tax + i.SubTotal)
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

package data

type Models struct {
	Invoices InvoiceModel
}

func NewModels() Models {
	return Models{
		Invoices: InvoiceModel{DB: make(map[string]Invoice)},
	}
}

package pdf

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Nico2220/billease/internal/data"
	"github.com/jung-kurt/gofpdf"
)

func New(inputData data.Invoice, from, to data.Company) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	// services := []data.Service{
	// 	{
	// 		ServiceType: "Develpment",
	// 		Description: "api developement",
	// 		Rate:        473,
	// 		Quantity:    1,
	// 		Amount:      473,
	// 	},

	// 	{
	// 		ServiceType: "Develpment",
	// 		Description: "api developement",
	// 		Rate:        473,
	// 		Quantity:    1,
	// 		Amount:      473,
	// 	},
	// }

	// inputData := data.InputData{
	// 	Services: services,
	// 	SubTotal: 50,
	// 	Tax:      20,
	// 	Total:    70,
	// 	Currency: "$",
	// 	Vat:      22,
	// }

	// from := data.Company{
	// 	Name:         "PhiTech",
	// 	Contact:      "Nicolas",
	// 	Adress:       "Germain street",
	// 	Country:      "Estonia",
	// 	SocityNumber: "123456",
	// 	Code:         "585943",
	// 	VatNumber:    "EE1234445",
	// 	PhoneNumber:  "+79772820353",
	// 	Email:        "nphilibert17@gmail.com",
	// }

	// to := data.Company{
	// 	Name:         "The Good Seat",
	// 	Contact:      "Alex",
	// 	Adress:       "Rue des entrepreneur",
	// 	Country:      "France",
	// 	SocityNumber: "123456",
	// 	Code:         "585943",
	// 	VatNumber:    "EE1234445",
	// 	PhoneNumber:  "+79772820353",
	// 	Email:        "alex@gmail.com",
	// }

	invoiceData := struct {
		From     data.Company
		To       data.Company
		Services []data.Service
		Vat      int
		Tax      float64
		Total    float64
		Currency string
		SubTotal float64
	}{
		From:     from, // Company
		To:       to,   // Company
		Services: inputData.Services,
		Vat:      inputData.Vat,
		Tax:      inputData.Tax,
		Total:    inputData.Total,
		Currency: inputData.Currency,
		SubTotal: inputData.SubTotal,
	}

	// invoiceData.CalculateSubTotal()
	// invoiceData.CalculateTax()
	// invoiceData.CalculateTotal()

	var opt gofpdf.ImageOptions
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(64, 64, 64)
	opt.ImageType = "png"

	pdf.ImageOptions("logo.png", -10, 10, 30, 0, false, opt, 0, "")

	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(64, 64, 64)
	pdf.CellFormat(0, 10, "Invoice", "0", 0, "R", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(64, 64, 64)
	pdf.Cell(40, 10, fmt.Sprintf("Date: %+v", time.Now().Format("Mon Jan 2, 2006 ")))

	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(64, 64, 64)
	pdf.CellFormat(0, 10, "Invoice No: 1234", "0", 0, "R", false, 0, "")

	pdf.Ln(12)

	//Invoice receipient info
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(255, 255, 250)
	pdf.CellFormat(80, 10, "Invoiced To:", "0", 0, "", false, 0, "")
	pdf.CellFormat(0, 10, "Pay To:", "0", 0, "R", false, 0, "")

	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 250)
	pdf.CellFormat(80, 10, invoiceData.From.Name, "0", 0, "", true, 0, "")
	pdf.CellFormat(0, 10, invoiceData.To.Name, "0", 0, "R", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 250)
	pdf.CellFormat(80, 10, invoiceData.From.Adress, "0", 0, "", true, 0, "")
	pdf.CellFormat(0, 10, invoiceData.To.Adress, "0", 0, "R", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 250)
	pdf.CellFormat(80, 10, "HP12 3JL", "0", 0, "", true, 0, "")
	pdf.CellFormat(0, 10, "Orange, CA 92865", "0", 0, "R", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 250)
	pdf.CellFormat(80, 10, "United Kingdom", "0", 0, "", true, 0, "")
	pdf.CellFormat(0, 10, "contact@koiceinc.com", "0", 0, "R", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(248, 249, 250)
	pdf.CellFormat(40, 10, "Service", "0", 0, "", true, 0, "")
	pdf.CellFormat(80, 10, "Description", "0", 0, "", true, 0, "")
	pdf.CellFormat(30, 10, "Rate", "0", 0, "", true, 0, "")
	pdf.CellFormat(20, 10, "QTY", "0", 0, "", true, 0, "")
	pdf.CellFormat(0, 10, "Amount", "0", 0, "R", true, 0, "")

	pdf.Ln(-1)

	for _, v := range invoiceData.Services {
		pdf.SetFont("Arial", "", 8)
		pdf.SetFillColor(255, 255, 255)
		pdf.CellFormat(40, 10, v.ServiceType, "0", 0, "L", false, 0, "")
		pdf.CellFormat(80, 10, v.Description, "0", 0, "L", false, 0, "")
		pdf.CellFormat(30, 10, strconv.FormatFloat(v.Rate, 'g', -1, 64), "0", 0, "L", false, 0, "")
		pdf.CellFormat(20, 10, strconv.FormatFloat(v.Quantity, 'g', -1, 64), "0", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, strconv.FormatFloat(v.Amount, 'g', -1, 64), "0", 0, "R", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(248, 249, 250)
	pdf.CellFormat(0, 10, fmt.Sprintf("Sub Total: %v", invoiceData.SubTotal), "0", 0, "R", true, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0, 10, fmt.Sprintf("Tax: %v", invoiceData.Tax), "0", 0, "R", true, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0, 10, fmt.Sprintf("Total: %v", invoiceData.Total), "0", 0, "R", true, 0, "")

	pdf.Ln(-1)

	if pdf.Err() {
		log.Fatal(pdf.Error())
	}

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println(err)
	}
}

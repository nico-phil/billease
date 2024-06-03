package data

type Company struct {
	ID           int64
	Name         string
	Contact      string
	Adress       string
	Country      string
	SocityNumber string
	Code         string
	VatNumber    string
	PhoneNumber  string
	Email        string
}

type CompanyModel struct {
	DB []Company
}

func (m *CompanyModel) Insert() (*Company, error) {
	return nil, nil
}

func (m CompanyModel) Get(id int64) (*Company, error) {
	return nil, nil
}

package application

import "github.com/evandrobarbosadosreis/go-rest-development/domain"

type CustomerReponse struct {
	Id      string `json:"id"`
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Birth   string `json:"date_of_birth"`
	Zipcode string `json:"zip_code"`
	Status  string `json:"status"`
}

func NewResponseFromDomain(c *domain.Customer) *CustomerReponse {
	return &CustomerReponse{
		Id:      c.Id,
		Name:    c.Name,
		City:    c.City,
		Birth:   c.Birth,
		Zipcode: c.Zipcode,
		Status:  domain.Status(c.Status).String(),
	}
}

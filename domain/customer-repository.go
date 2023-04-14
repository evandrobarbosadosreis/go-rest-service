package domain

type CustomerRepository interface {
	FindAll(status Status) ([]Customer, *AppError)
	FindById(id string) (*Customer, *AppError)
}




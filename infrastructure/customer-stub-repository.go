package infrastructure

import "github.com/evandrobarbosadosreis/go-rest-development/domain"

type customerRepositoryStub struct {
	customers []domain.Customer
}

func (repository *customerRepositoryStub) FindAll(status domain.Status) ([]domain.Customer, *domain.AppError) {
	return repository.customers, nil
}

func (*customerRepositoryStub) FindById(id string) (*domain.Customer, *domain.AppError) {
	result := domain.Customer{
		Id:   id,
		Name: "Exemple",
		City: "Somewhere",
	}
	return &result, nil
}

func NewStubRepository() domain.CustomerRepository {
	customers := []domain.Customer{
		{Id: "1001", Name: "John", City: "Something"},
		{Id: "1002", Name: "Anna", City: "Someplace"},
		{Id: "1003", Name: "Eddy", City: "Somewhere"},
		{Id: "1004", Name: "Brad", City: "Sometime"},
	}
	return &customerRepositoryStub{customers: customers}
}

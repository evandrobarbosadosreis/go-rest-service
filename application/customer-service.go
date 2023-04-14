package application

import (
	"github.com/evandrobarbosadosreis/go-rest-development/domain"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *domain.AppError)
	GetCustomer(id string) (*domain.Customer, *domain.AppError)
}

type defaultCustomerService struct {
	repository domain.CustomerRepository
}

func (service *defaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *domain.AppError) {
	return service.repository.FindAll(domain.Status(status))
}

func (service *defaultCustomerService) GetCustomer(id string) (*domain.Customer, *domain.AppError) {
	return service.repository.FindById(id)
}

func NewDefaultCustomerService(repository domain.CustomerRepository) CustomerService {
	return &defaultCustomerService{repository: repository}
}

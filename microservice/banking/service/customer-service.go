package service

import (
	"github.com/banking/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	CustomerRepository domain.ICustomerRepository
}

func NewCustomerService(repository domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{CustomerRepository: repository}
}

func (c DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return c.CustomerRepository.FindAll()
}

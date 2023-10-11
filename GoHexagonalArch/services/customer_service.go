package services

import (
	"database/sql"
	"goarch/errs"
	"goarch/logs"
	"goarch/repositories"
)

type customerService struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) CustomerService {
	return customerService{customerRepo: customerRepo}
}
func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers := []CustomerResponse{}
	customersFromRepo, err := s.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	for _, customer := range customersFromRepo {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customers = append(customers, customerResponse)
	}
	return customers, nil
}
func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer := CustomerResponse{}
	customerFromRepo, err := s.customerRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customer = CustomerResponse{
		CustomerID: customerFromRepo.CustomerID,
		Name:       customerFromRepo.Name,
		Status:     customerFromRepo.Status,
	}
	return &customer, nil
}

package repositories

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1, Name: "John", DathOfBirth: "2000-01-01", City: "New York", ZipCode: "10000", Status: 1},
		{CustomerID: 2, Name: "Smith", DathOfBirth: "2000-01-01", City: "New York", ZipCode: "10000", Status: 1},
		{CustomerID: 3, Name: "Ben", DathOfBirth: "2000-01-01", City: "New York", ZipCode: "10000", Status: 1},
		{CustomerID: 4, Name: "Peter", DathOfBirth: "2000-01-01", City: "New York", ZipCode: "10000", Status: 1},
		{CustomerID: 5, Name: "Jane", DathOfBirth: "2000-01-01", City: "New York", ZipCode: "10000", Status: 1},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}

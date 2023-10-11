package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Test", 18, nil)
	c.On("GetCustomer", 2).Return("Tes2t", 18, errors.New("not found"))

	name, age, err := c.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("name : %v\nage : %v\n", name, age)

}

type CustomerRepository interface {
	GetCustomer(id int) (name string, age int, err error)
	Hello()
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}

func (m *CustomerRepositoryMock) Hello() {
	println(m)
}

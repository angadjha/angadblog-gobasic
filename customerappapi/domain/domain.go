package domain

import "errors"

var (
	ErrNotFound       = errors.New("no records found")
	ErrCustomerExists = errors.New("customer exists")
)

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerRepository interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}

package memstore

import (
	"customerappapi/domain"
	"errors"
	"fmt"
)

// CustomerRepository provides memory based local data repository.
type CustomerRepository struct {
	repository map[string]domain.Customer
}

func NewCustomerRepository() (*CustomerRepository, error) {
	return &CustomerRepository{
		repository: make(map[string]domain.Customer),
	}, nil
}

// Create adds a new customer to the repository.
func (cr *CustomerRepository) Create(c domain.Customer) error {
	if _, ok := cr.repository[c.ID]; ok {
		return errors.New("customer already exists")
	}
	cr.repository[c.ID] = c
	return nil
}

// Update updates a customer in the repository.
func (cr *CustomerRepository) Update(id string, c domain.Customer) error {
	if _, ok := cr.repository[id]; !ok {
		return fmt.Errorf("customer id ( %v ) not found", id)
	}
	cr.repository[id] = c
	return nil
}

// Delete removes a customer from the repository.
func (cr *CustomerRepository) Delete(id string) error {
	if _, ok := cr.repository[id]; !ok {
		return fmt.Errorf("customer id ( %v ) not found", id)
	}
	delete(cr.repository, id)
	return nil
}

// GetById retrieves a customer from the repository by ID.
func (cr *CustomerRepository) GetById(id string) (domain.Customer, error) {
	c, ok := cr.repository[id]
	if !ok {
		return domain.Customer{}, fmt.Errorf("customer id ( %v ) not found", id)
	}
	return c, nil
}

// GetAll retrieves all customers from the repository.
func (cr *CustomerRepository) GetAll() ([]domain.Customer, error) {
	if len(cr.repository) == 0 {
		return nil, errors.New("no data")
	}
	var customers []domain.Customer
	customers = make([]domain.Customer, 0, len(cr.repository))
	for _, c := range cr.repository {
		customers = append(customers, c)
	}
	return customers, nil
}

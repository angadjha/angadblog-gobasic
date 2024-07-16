package main

import (
	"bytes"
	"customerappapi/domain"
	"customerappapi/memstore"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	// Initialize a new in-memory repository
	repo, _ := memstore.NewCustomerRepository()
	// Initialize the customer controller with the repository
	customerController := CustomerHandler{repository: repo}

	// Create a new customer
	newCustomer := domain.Customer{
		ID:    "1",
		Name:  "Angad",
		Email: "angad@gmail.com",
	}

	// Marshal the customer struct to JSON
	requestBody, err := json.Marshal(newCustomer)
	if err != nil {
		t.Errorf("Failed to marshal customer to JSON: %v", err)
	}

	// Create a new HTTP POST request with the JSON payload
	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(customerController.Post)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check if the customer was added to the repository
	_, err = repo.GetById("1")
	if err != nil {
		t.Errorf("customer not added to repository: %v", err)
	}
}

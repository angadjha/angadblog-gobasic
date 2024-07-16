package main

import (
	"customerappapi/domain"
	"encoding/json"
	"errors"
	"net/http"
)

// CustomerHandler Organises the CRUD operations at UI layer.
type CustomerHandler struct {
	repository domain.CustomerRepository
}

// Post handles HTTP Post - /api/customer
func (ch *CustomerHandler) Post(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	// Decode the incoming customer json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create Customer
	if err := ch.repository.Create(customer); err != nil {
		if errors.Is(err, domain.ErrCustomerExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAll handles HTTP Get - /api/customer
func (ch *CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// Get all
	if customers, err := ch.repository.GetAll(); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		j, err := json.Marshal(customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}

// Get handles HTTP Get - /api/customer/{id}
func (ch *CustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Getting route parameter id
	id := r.PathValue("id")
	// Get by id
	if customer, err := ch.repository.GetById(id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// Put handles HTTP Put - /api/customer/{id}
func (ch *CustomerHandler) Put(w http.ResponseWriter, r *http.Request) {
	// Getting route parameter id
	id := r.PathValue("id")
	var customer domain.Customer
	// Decode the incoming customer json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update
	if err := ch.repository.Update(id, customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Delete handles HTTP Delete - /api/customer/{id}
func (ch *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Getting route parameter id
	id := r.PathValue("id")
	// delete
	if err := ch.repository.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

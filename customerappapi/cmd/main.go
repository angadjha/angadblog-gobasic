package main

import (
	"customerappapi/memstore"
	"log"
	"net/http"
)

func main() {
	repo, err := memstore.NewCustomerRepository() // With in-memory database
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &CustomerHandler{
		repository: repo, // Injecting dependency
	}
	router := initializeRoutes(h) // configure routes

	server := &http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	log.Println("Listening...")
	err = server.ListenAndServe()
	if err != nil {
		return
	} // Run the http server
}

func initializeRoutes(h *CustomerHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/customers", h.GetAll)
	mux.HandleFunc("GET /api/customers/{id}", h.Get)
	mux.HandleFunc("POST /api/customers", h.Post)
	mux.HandleFunc("PUT /api/customers/{id}", h.Put)
	mux.HandleFunc("DELETE /api/customers/{id}", h.Delete)
	return mux
}

package main

import (
	"net/http"

	"github.com/banking/domain"
	"github.com/banking/service"

	"github.com/gorilla/mux"
)

func StartServer() {

	router := mux.NewRouter()
	ch := &CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", ch.getCustomerWithId).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.createCustomer).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)
}

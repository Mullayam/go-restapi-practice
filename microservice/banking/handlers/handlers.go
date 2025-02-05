package main

import (
	"encoding/json"
	"net/http"

	"github.com/banking/service"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	id   int    `json:"id" xml:"id"`
	Job  string `json:"job" xml:"job"`
}
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
func (ch *CustomerHandler) getCustomerWithId(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["customerId"]

	Customer := Customer{Name: "John Doe", id: 30, Job: "Software Engineer"}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Customer)

}
func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

}

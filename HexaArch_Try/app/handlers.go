package app

import (
	"HexaArch_Try/service"
	"encoding/json"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	json.NewEncoder(w).Encode(customers)
}

package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name   string `json:"product_name"` // name to be displayed in the response i.e json --> option to provide alias name or lowercase var names in the response
	Id     int    // produces caps only in the beginning of the var name
	Amount int
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Laptop", Id: 12458, Amount: 50000},
		{Name: "Tablet", Id: 25874, Amount: 15000},
	}
	w.Header().Add("Content-Type", "application/json") // without this statement, the header in the postman tab shows content-type as plain text
	json.NewEncoder(w).Encode(products)
}

func main() {

	http.HandleFunc("/getproducts", getProducts)
	http.ListenAndServe("localhost:8085", nil)
}

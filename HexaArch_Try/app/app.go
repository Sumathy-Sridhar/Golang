package app

import (
	"HexaArch_Try/domain"
	"HexaArch_Try/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

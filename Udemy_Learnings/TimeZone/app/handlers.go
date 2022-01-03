package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	router.HandleFunc("/api/time", getTimeZone).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

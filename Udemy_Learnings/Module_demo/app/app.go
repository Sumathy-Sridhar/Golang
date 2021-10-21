package app

import (
	"net/http"
)

func Start() {
	// http.HandleFunc("/getemp", getAllEmployees)
	// http.ListenAndServe("localhost:8000", nil)

	//Refractoring of http
	Mux := http.NewServeMux()
	Mux.HandleFunc("/getemp", getAllEmployees)
	http.ListenAndServe("localhost:8000", Mux)
}

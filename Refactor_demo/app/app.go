package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/getcos", getCosmetics)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

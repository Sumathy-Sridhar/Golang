package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", greet) // defeines the routes

	//http.ListenAndServe("localhost:8082", nil) // starting the server
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Worold Program From Udemy!!")
}

package main

import (
	"Video/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("server started on port :: 5000")
	router := mux.NewRouter()

	router.HandleFunc("/api", api.UploadFiles).Methods("Post")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println(err)
	}

}

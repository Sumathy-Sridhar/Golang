package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()
	//router.HandleFunc("/getcos", getCosmetics)

	// trying to send the id along with the url to display in the postman
	//router.HandleFunc("/getcos/{cos_id}", getcosid)

	//By default the http request is get method so we dont mention it explicity, if in case to mention tat we can do this way
	router.HandleFunc("/getcos", getCosmetics).Methods(http.MethodGet)

	// this strictly allows only numeric values in the url for value of id if not throws 404 error
	router.HandleFunc("/getcos/{cos_id:[0-9]+}", getcosid).Methods(http.MethodGet)

	//Post Method
	router.HandleFunc("/posturl", postdemo).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8001", router))
	//Make all the handler functions before the log.fatal statement. Any request made after the log.fatal statement returns only 404
}

// function to accept the id from the url and display in the postman
func getcosid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["cos_id"])
}

//Post Function
func postdemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post Request Demo...")
}

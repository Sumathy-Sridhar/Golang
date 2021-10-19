package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Socialmedia struct {
	Appname     string `json:"Appname"`
	Noofusers   int    `json:"Noofusers"`
	Description string `json:"Description"`
}

var media []Socialmedia

func demopage(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "RestAPi Demo")
	fmt.Println("Response Rendered successfully ")
}
func allsocialmedia(so http.ResponseWriter, rq *http.Request) {
	media = []Socialmedia{
		{Appname: "Instagram", Noofusers: 100000, Description: "Messaging and memes app"},
		{Appname: "Whatsapp", Noofusers: 5000, Description: "Messaging app"},
		{Appname: "Youtube", Noofusers: 125000, Description: "Entertainment and Fun App"},
	}
	fmt.Println("Social Media contents")
	json.NewEncoder(so).Encode(media)
}
func handleRequest() {
	http.HandleFunc("/", demopage)
	http.HandleFunc("/somedia", allsocialmedia)
	//http.HandleFunc("/postmedia")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}

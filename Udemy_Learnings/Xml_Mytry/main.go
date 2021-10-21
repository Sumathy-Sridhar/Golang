package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

type Cosmetics struct {
	Name    string `json:"Cosmetic_name" xml:"cos_name"`
	Purpose string `json:"Usage" xml:"purpose"`
	Amount  int    `json:"amt" xml:"Amt"`
}

func getCosmetics(w http.ResponseWriter, r *http.Request) {
	cosmetics := []Cosmetics{
		{Name: "Lipstick", Purpose: "Enhance lips", Amount: 500},
		{Name: "Eyeliner", Purpose: "Sharpen eyes", Amount: 600},
		{Name: "Mascara", Purpose: "Brighten eyelashes", Amount: 700},
	}
	// w.Header().Add("Content-type", "application/xml")
	// xml.NewEncoder(w).Encode(cosmetics)

	/* if we set the content-type in the header tab (postman), this part of code tries to match the content-type
	given in the headers and does the encoding based on the type*/
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(cosmetics)
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(cosmetics)
	}
}
func main() {
	http.HandleFunc("/getcos", getCosmetics)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

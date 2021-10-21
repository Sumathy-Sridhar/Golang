package app

import (
	"encoding/xml"
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
	w.Header().Add("Content-type", "application/xml")
	xml.NewEncoder(w).Encode(cosmetics)

}

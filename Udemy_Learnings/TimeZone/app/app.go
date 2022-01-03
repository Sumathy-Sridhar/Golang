package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func getTimeZone(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) <= 1 {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
		} else {
			response["current_time"] = time.Now().In(loc).String()
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		for _, tzdb := range timezones {
			loc, err := time.LoadLocation(tzdb)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezone %s in input", tzdb)))
				return
			}
			now := time.Now().In(loc)
			response[tzdb] = now.String()
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

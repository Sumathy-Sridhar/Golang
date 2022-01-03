package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Username string
	Password string
	// Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("Login.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Username: r.FormValue("Username"),
			Password: r.FormValue("Password"),
			// Message: r.FormValue("message"),
		}

		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})

		fmt.Println("login details::", details)
	})

	http.ListenAndServe(":8080", nil)
}

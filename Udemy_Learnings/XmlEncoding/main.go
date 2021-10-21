package main

import (
	"encoding/xml"
	"net/http"
)

type Employee struct {
	EmployeeName string
	EmployeeId   int `json:"employee_id" xml:"employeeid"`
	// Both json and xml format for var names can be given in the same struct no need of creating 2 diff structs bcoz based on the type of encoding as json/xml it produces the response
	Designation string `xml:"employee_designation"`
	Sal         int    `json:"Salary"`
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		{EmployeeName: "Sumathy", EmployeeId: 12345, Designation: "Developer", Sal: 50000},
		{EmployeeName: "Delvin", EmployeeId: 54321, Designation: "Tester", Sal: 40000},
		{EmployeeName: "Mathew", EmployeeId: 25874, Designation: "Business Analyst", Sal: 450000},
	}
	w.Header().Add("Content-Type", "application/xml")
	//json.NewEncoder(w).Encode(employees)
	xml.NewEncoder(w).Encode(employees)
}
func main() {
	http.HandleFunc("/getemp", getAllEmployees)
	http.ListenAndServe("localhost:8000", nil)
}

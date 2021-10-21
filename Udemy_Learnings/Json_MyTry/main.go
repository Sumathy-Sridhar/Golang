package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	EmployeeName string
	EmployeeId   int `json:"employee_id"`
	Designation  string
	Sal          int `json:"Salary"`
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		{EmployeeName: "Sumathy", EmployeeId: 12345, Designation: "Developer", Sal: 50000},
		{EmployeeName: "Delvin", EmployeeId: 54321, Designation: "Tester", Sal: 40000},
		{EmployeeName: "Mathew", EmployeeId: 25874, Designation: "Business Analyst", Sal: 450000},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
func main() {
	http.HandleFunc("/getemp", getAllEmployees)
	http.ListenAndServe("localhost:8000", nil)
}

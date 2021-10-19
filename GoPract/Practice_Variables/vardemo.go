package main

import "fmt"

func main() {
	var id int
	fmt.Println("Empid: ", id)
	id = 123
	fmt.Println("Empid: ", id)
	var name = "Sumathy"
	fmt.Println("Employee Name: ", name)
	var salary, experience int
	salary = 10000
	experience = 2
	fmt.Println("Salary: ", salary)
	fmt.Println("Experience: ", experience)
	var designation, department = "Trainee", "CCx"
	fmt.Println("Designation: ", designation)
	fmt.Println("Department: ", department)
	a, b := 10, 20
	fmt.Println(a)
	fmt.Println(b)
	var (
		empname = "Delvin"
		empid   = 143
		sal     int
	)
	fmt.Println("Empname: ", empname, " Empid: ", empid, "Salary", sal)

	age, technology := 24, "Golang"
	fmt.Println("Age: ", age)
	fmt.Println("Technology: ", technology)
}

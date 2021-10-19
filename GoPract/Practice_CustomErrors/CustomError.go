package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	empid       int
	empname     string
	designation string
}

func details(emp Employee) error {
	if emp.empid <= 0 {
		return errors.New("Employee Id cannot be Zero or Negative!")
	}
	if emp.empname == "" {
		return errors.New("Employee Name cannot be Empty...")
	}
	if emp.designation == "" {
		return errors.New("Desgination is Mandatory!")
	}
	return nil
}

func main() {
	error := details(Employee{empid: -3, empname: "Sumathy", designation: "Developer"})
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("Employee data is verified...")
}

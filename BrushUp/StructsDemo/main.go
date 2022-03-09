package main

import "fmt"

//Structs are mutable.
type employee struct {
	name  string
	empid int
}

func newemp(name string) *employee {
	e := employee{name: name}
	e.empid = 10
	return &e
}
func main() {
	fmt.Println(employee{"Sumathy", 36})
	fmt.Println(employee{empid: 12})
	fmt.Println(newemp("Darling"))
}

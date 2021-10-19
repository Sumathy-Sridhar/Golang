package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {
	emp1 := employee{
		salary:  12000,
		country: "USA",
	}
	emp2 := employee{
		salary:  14000,
		country: "Canada",
	}
	emp3 := employee{
		salary:  13000,
		country: "India",
	}
	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	}
	fmt.Println("Length of Map: ", len(employeeInfo))
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)
	// map1 := map[string]int{
	// 	"one": 1,
	// 	"two": 2,
	// }

	// map2 := map1

	// if map1 == map2 { //map can only be compared to nil to check whether two maps are equal is to compare each one's individual elements one by one or using reflection
	// }
}

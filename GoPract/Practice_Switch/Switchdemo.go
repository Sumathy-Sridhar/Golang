package main

import "fmt"

func multiply() int {
	num1 := 15 * 5
	return num1
}
func main() {
	var num int
	fmt.Println("Enter number: ")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("Male")
	case 2:
		fmt.Println("FeMale")
	case 3:
		fmt.Println("Transgender")
	// case 3:
	// 	fmt.Println("Transgender")
	default:
		fmt.Println("Invalid number")
	}
	var no int
	fmt.Println("Enter number to check no is +ve,-ve or zero: ")
	fmt.Scanln(&no)
	switch {
	case no > 0:
		fmt.Println("Positive")
	case no < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
	switch numm := multiply(); { //num is not a constant
	case numm < 50:
		fmt.Printf("%d is lesser than 50\n", numm)
		fallthrough
	case numm < 100:
		fmt.Printf("%d is lesser than 100\n", numm)
	}
}

package main

import "fmt"

type add func(a int, b int) int //Userdefined Types

func highodr(num1 func(num1, num2 int) int) {
	fmt.Println(num1(10, 20))
}

func main() {
	a := func() { // annoymous fn (without name)
		fmt.Println("Anonymous function assinged to varaible a")
	}
	a() // calling of func since it is assigned to a
	fmt.Printf("Type of a is : %T", a)
	fmt.Println()
	func() { //annoymous fn (without name)
		fmt.Println("Annonymous function not assigned to any variable")
	}() // calling func
	func(name string) { // passing parameters to annonymous fn
		fmt.Println("Hi", name)
	}("Sumathy!") // calling of fn with parameter

	//Userdefined Types
	var ad add = func(a, b int) int {
		return a + b
	}
	sum := ad(10, 20)
	fmt.Println("Sum=", sum)
	f := func(num1, num2 int) int {
		return num1 * num2
	}
	highodr(f)
}

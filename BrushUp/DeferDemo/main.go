package main

import "fmt"

func dummy() {
	fmt.Println("Dummy Function..")
}

func printNum(a int) {
	fmt.Println(a)
}

// defer stat executes at the end when its in the main func
func main() {

	// multiple defer statements are stacked and executed in the order of LIFO
	defer dummy()
	defer printNum(1)
	defer printNum(2)
	defer printNum(3)
	defer fmt.Println("Defer statement...")
	fmt.Println("Main Function..")
}

package main

import (
	"fmt"
)

func printA(a int) {
	fmt.Println()
	fmt.Println("value of a in deferred function", a)
}
func main() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
	name := "Sumathy"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v) //When a function has multiple defer calls, they are pushed on to a stack and executed in Last In First Out (LIFO) order.
	}
}

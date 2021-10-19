package main

import "fmt"

func add(numbers ...int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func main() {
	result := add(1, 2, 3, 4)
	fmt.Println("Sum: ", result)
	sumres := add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Sum of first 10 numbers: ", sumres)
}

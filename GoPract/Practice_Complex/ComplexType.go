package main

import "fmt"

func main() {

	c1 := complex(2, 5)
	c2 := 2 + 3i
	res := c1 + c2
	fmt.Println("Addition Result: ", res)
	mulres := c1 * c2
	fmt.Println("Multiplication Result: ", mulres)
}

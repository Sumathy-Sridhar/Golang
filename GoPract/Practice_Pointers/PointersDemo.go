package main

import "fmt"

func main() {
	val := 255
	var value *int = &val
	fmt.Printf("Type of value is %T\n", value)
	fmt.Println("address of val is", value)
	// zero value pointer
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}

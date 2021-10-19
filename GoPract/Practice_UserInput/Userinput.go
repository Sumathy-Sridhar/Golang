package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter a and b: ")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	c := a + b
	fmt.Printf("%d + %d: %d\n", a, b, c)
	fmt.Println("Dinesh")
}

package main

import "fmt"

func main() {
	first := "Sumathy"
	last := "Sridhar"
	name := first + " " + last
	fmt.Println("My name is", name)

	//Type conversion
	num := 10
	floatnum := 23.3
	sum := num + int(floatnum) //explicit conversion
	fmt.Println(sum)

	//diff bet constants and var
	var a = 10
	fmt.Println(a)
	a = 30
	fmt.Println(a)
	const c = 20
	fmt.Println(c)
	//c = 40  reassignment is not allowed with constants
	var ab = 5.9 / 8
	fmt.Printf("ab's type is %T and value is %v", ab, ab)

}

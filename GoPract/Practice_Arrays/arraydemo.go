package main

import "fmt"

func main() {
	var b [5]string
	fmt.Println(b)
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
	ar := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(ar)
	arr := [...]int{1} // ... makes the compiler determine the length
	fmt.Println(arr)
	strarr := [...]string{"Brown", "cony"}
	sarr := strarr
	sarr[1] = "Delvin"
	fmt.Println("Original Array: ", strarr)
	fmt.Println("Copy of Array: ", sarr) // arrays are of value types and not reference types so they any chges made in the new variable will not be reflected in the original array
}

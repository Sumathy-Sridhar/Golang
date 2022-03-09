package main

import "fmt"

func main() {

	//array decl
	var a [3]int
	fmt.Println(a)

	a[2] = 10
	fmt.Println("Set array: ", a)
	fmt.Println("Get a[2]: ", a[2])
	fmt.Println("Length of a: ", len(a))

	//shorthand declaration
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	//2D array
	var twoD [10][5]int
	for i := 5; i < 10; i++ {
		for j := 1; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}

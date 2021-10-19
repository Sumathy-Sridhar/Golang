package main

import "fmt"

func printarray(a [3][2]int) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Printf("\n")
	}
}
func main() {
	arr := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	printarray(arr)
}

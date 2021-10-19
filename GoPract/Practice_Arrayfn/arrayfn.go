package main

import "fmt"

func changeLocal(num [5]float64) {
	num[2] = 45.25
	fmt.Println("inside function ", num)

}
func main() {
	num := [...]float64{25.3, 70.00, 26.22, 82.03, 33.55}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
	fmt.Println("Length of Array: ", len(num))
	fmt.Println("Elements of Array iterating using for loop: ")
	for i := 0; i < len(num); i++ {
		fmt.Printf("Element %d = %.2f\n", i, num[i])
	}

}

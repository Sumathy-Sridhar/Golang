package main

import "fmt"

func add(no1, no2 int) int {
	var result = no1 + no2
	return result
}

// func rectdime(len, bred float64) (float64, float64) { //2 return types since it returns 2 values
// 	var area = len * bred
// 	var permi = 2 * (len + bred)
// 	return area, permi
// }

//named return values
func rectdime(len, bred float64) (area, permi float64) { //2 return types since it returns 2 values
	area = len * bred
	permi = 2 * (len + bred)
	return area, permi
}
func main() {
	no1, no2 := 10, 20
	result := add(no1, no2)
	fmt.Println("Addition result: ", result)
	// area, permi := rectdime(10.5, 20.4)
	// fmt.Println("Area: ", area)
	// fmt.Println("Perimeter: ", permi)
	area, _ := rectdime(20.5, 30.4) //blank identifier discards the perimeter
	fmt.Println("Area: ", area)
}

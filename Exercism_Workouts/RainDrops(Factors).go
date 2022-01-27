package main

import (
	"fmt"
	"strconv"
)

var str string
var number int

func Convert(number int) {

	if number%3 == 0 {
		str += "Pling"
	}
	if number%5 == 0 {
		str += "Plang"
	}
	if number%7 == 0 {
		str += "Plong"
	}
	if str == "" {
		str = strconv.Itoa(number)
	}
	fmt.Println(str)
}
func main() {
	fmt.Println("Enter number to check Rainndrops: ")
	fmt.Scanf("%d", &number)
	Convert(number)
}

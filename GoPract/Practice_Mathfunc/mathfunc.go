package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 143.2, 25.25
	c := math.Min(a, b)
	fmt.Println("Minimum : ", c)
	d := math.Max(a, b)
	fmt.Println("Maximum: ", d)
	e := math.Round(125.75)
	fmt.Println("Round off value of", e)
}

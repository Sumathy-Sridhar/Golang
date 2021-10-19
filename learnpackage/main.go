package main

import (
	"fmt"
	"learnpackage/simpleinterest"
)

func main() {
	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := simpleinterest.Calculate(p, r, t) // func names defined with first letter caps can only be exported and accessed outside the package else compiler throws error
	fmt.Println("Simple interest is", si)
}

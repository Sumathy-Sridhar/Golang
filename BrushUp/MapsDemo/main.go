package main

import "fmt"

func main() {

	mp := make(map[string]int)

	mp["a"] = 10
	mp["b"] = 20

	fmt.Println(mp)
	val := mp["b"]
	fmt.Println("Value of b: ", val)

	mpint := make(map[int]int)

	mpint[01] = 10
	mpint[02] = 20
	mpint[3] = 20

	fmt.Println(mpint)
	fmt.Println("Length: ", len(mpint))
	delete(mpint, 2)
	fmt.Println(mpint)
}

package main

import "fmt"

func appendstr() func(string) string {
	t := "Hello"
	c := func(str string) string {
		t = t + " " + str
		return t
	}
	return c
}

func main() {
	a := appendstr()
	b := appendstr()
	fmt.Println(a("Sumathy!"))
	fmt.Println(b("Everyone..."))

	fmt.Println(a("Welcome..."))
	fmt.Println(b("Have a Nice Day!!"))
}

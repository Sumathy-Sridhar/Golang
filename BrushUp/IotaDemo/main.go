package main

import "fmt"

/*iota - AutoIncrement of the value
works only for constants
value of the constant starts with 0
increase by 1 after each line */

func main() {

	const (
		a = iota     // 0
		b = iota + 4 // 1+4
		c = iota * 5 // 2*5
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

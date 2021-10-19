package main

import (
	"fmt"
	"unsafe" // indicates non-portable
)

func main() {
	var a int = 89
	b := "Sumathy"
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b))
}

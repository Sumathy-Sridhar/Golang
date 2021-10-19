package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("D:\\Demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "Opened Successfully")
}

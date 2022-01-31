package main

import "fmt"

func ShareWith(name string) {
	if name == "" {
		name = "you"
	}
	fmt.Println("One for " + name + ", one for me.")
}

func main() {
	ShareWith("Sumathy")
	ShareWith("")
}

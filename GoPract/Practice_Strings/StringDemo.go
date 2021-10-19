package main

import "fmt"

//Accessing individual bytes of a string
func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // %x returns the hexadecimal value
	}
}

//Accessing individual characters of a string
func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	var name string
	fmt.Println("Enter string: ")
	fmt.Scanln(&name)
	fmt.Printf("Entered String : %s\n", name)
	printChars(name) // returns the characters of the entered string
	fmt.Println()
	printBytes(name) // returs unicode UTF-8 encoded values of given string
}

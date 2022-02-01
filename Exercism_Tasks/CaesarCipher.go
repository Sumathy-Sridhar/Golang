package main

import (
	"fmt"
	"unicode"
)

// RotationalCipher is an implementation for the rotational cipher exercise.
func RotationalCipher(plain string, shiftKey int) {
	var cipher []rune
	for _, ru := range plain {
		if unicode.IsUpper(ru) {
			ru = rune('A' + (int(ru-'A')+shiftKey)%26)
		} else if unicode.IsLower(ru) {
			ru = rune('a' + (int(ru-'a')+shiftKey)%26)
		}
		cipher = append(cipher, ru)
	}

	fmt.Println(string(cipher))
}

func main() {
	RotationalCipher("Sumathy", 5)
}

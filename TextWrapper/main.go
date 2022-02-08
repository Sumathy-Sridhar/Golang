package main

import (
	"fmt"
	"unicode"
)

func main() {
	//const text = `Go is a statically typed, compiled programming language designed at Google by Robert Griesemer,Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. `

	const text = `It is a condition in which each case does not have any break statement. Note that in a switch statement, it is not necessary to have a break statement. Using the break statement is optional but using a break statement is a good habit. In this situation, the control flow will fall through until a break statement is found. It is the rarest situation.

	The break statement denotes the end of the case. Omitting a break leads to a program execution continuing into the next case and onwards till a break statement is encountered or the end of the switch is reached.`
	const maxWidth = 40

	var lw int // line width

	for _, r := range text {
		fmt.Printf("%c", r)

		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
	fmt.Println()
}

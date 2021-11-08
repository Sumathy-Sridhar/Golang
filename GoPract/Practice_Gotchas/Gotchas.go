package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {

	// *Invalid operation: mismatched types int and float64*
	// var x, y = 13, 3.5
	// fmt.Println(x / y)
	var x = 13 / 3.5 //he right side of = are constants, not variables. Hence compiler will convert 13 into a float type to execute the program.
	fmt.Println(x)
	var x1, y1 = 13, 3.5
	fmt.Println(float64(x1) / y1) //Use float64 to convert type of x into float64.

	//panic: assignment to entry in nil map
	// var rect map[string]int // Map types are reference types, like pointers or slices, and so the value of rect is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that.
	// rect["height"] = 10
	// fmt.Println(rect["height"])

	//recovering panic method 1
	var recta map[string]int
	fmt.Println(recta["height"]) //The Zero Value of an uninitialized map is nil. Both len and accessing the value of rect["height"] will work on nil map. len returns 0 and the key of "height" is not found in map and you will get back zero value for int which is 0. Similarly, idx will return 0 and key will return false.
	fmt.Println(len(recta))

	idx, key := recta["height"]
	fmt.Println(idx)
	fmt.Println(key)

	//recovering panic method 2
	var rectaw = map[string]int{"height": 10}
	fmt.Println(rectaw["height"])

	s := `Go\tJava\nPython` //Raw strings are enclosed in back-ticks `. Here, \t and \n has no special meaning, they are considered as backslash with t and backslash with n.
	fmt.Println(s)

	//If you need to include backslashes, double quotes or newlines in your string, use a raw string literal.
	str := "Go\tJava\nPython"
	fmt.Println(str)

	//invalid operation: mismatched types int and time.Duration
	// var timeout = 3
	// fmt.Println(timeout)
	// fmt.Println(timeout * time.Millisecond) //The operands to numeric operations must have the same type unless the operation involves shifts or untyped constants.

	const timeout = 10
	fmt.Println(timeout)
	fmt.Println(timeout * time.Millisecond) //Millisecond's underlying type is an int64, which the compiler knows how to convert to. Literals and constants are untyped until they are used, unless the type is explicitly declared. timeout is an untyped constant in this example. Its type is implicitly converted to time.Millisecond

	// returns the length of string , you will get the size in bytes in go
	data := "We♥Go"
	fmt.Println(len(data)) //In Go Strings are UTF-8 encoded, this means each charcter called rune can be of 1 to 4 bytes long. Here,the charcter ♥ is taking 3 bytes, hence the total length of string is 7.
	fmt.Println(utf8.RuneCountInString(data))

	// cannot use nil (untyped nil value) as string value in variable declarationcompilerIncompatibleAssign
	// var ata string = nil
	// if ata == nil {
	// 	fmt.Println(ata)
	// }

	var ata *string = nil //*string is the type of the pointer variable which points to a value of type string. The zero value of a pointer is nil.
	if ata == nil {
		fmt.Println(ata)
	}

	//floating point multiplication
	var m = 1.39
	fmt.Println(m * m)

	const n = 1.39
	fmt.Println(n * n)

	// Go doesn't allow automatic type promotion between variables. string type conversion
	// i := 105
	// s := string(i) //no new variables on left side of :=
	// fmt.Println(s)

	const i = 100
	//var i1 = 100 //i1 declared but not used
}

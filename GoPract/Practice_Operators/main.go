package main

import "fmt"

func main() {
	//Arithmetic operators
	var a, b = 20, 10
	fmt.Printf("a+b= %d\n", a+b)
	fmt.Printf("a-b= %d\n", a-b)
	fmt.Printf("a*b= %d\n", a*b)
	fmt.Printf("a mod b= %d\n", a%b)
	fmt.Printf("a/b= %d\n", a/b)

	//Assignment operators
	var x, y = 100, 200
	x = y
	fmt.Println("Assign (=) ", x)
	x = 15
	x += y
	fmt.Println("Add and Assign (+=) ", x)
	x = 40
	x -= y
	fmt.Println("Subtract and Assign (-=)", x)
	x = 50
	x *= y
	fmt.Println("Multiply and Assign (*=)", x)
	x = 10
	x /= y
	fmt.Println("Divide and Assign (/=)", x)
	x = 40
	x %= y
	fmt.Println("Mod and Assign (%=)", x)

	//Comparison operators
	var num1, num2 = 10, 20
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 <= num2)
	fmt.Println(num1 >= num2)

	//Logical operators
	var n1, n2, n3 = 50, 60, 70
	fmt.Println(n1 > n2 && n2 > n3)
	fmt.Println(n1 > n2 || n2 < n3)
	fmt.Println(!(n1 > n2 && n2 > n3))

	//Bitwise operators

	var x1 uint = 9  //0000 1001
	var y1 uint = 65 //0100 0001
	var z1 uint

	z1 = x1 & y1 //Sets each bit to 1 if both bits are 1
	fmt.Println("x & y  =", z1)

	z1 = x1 | y1 // Sets each bit to 1 if one of two bits is 1
	fmt.Println("x | y  =", z1)

	z1 = x1 ^ y1 // Sets each bit to 1 if only one of two bits is 1
	fmt.Println("x ^ y  =", z1)

	z1 = x1 << 1 // Shift left by pushing zeros in from the right and let the leftmost bits fall off
	fmt.Println("x << 1 =", z1)

	z1 = x1 >> 1 // Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off
	fmt.Println("x >> 1 =", z1)
}

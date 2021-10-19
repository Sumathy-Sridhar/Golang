package main

import "fmt"

// Interfaces with common methods
type shape interface {
	area() int
	perimeter() int
}

type square struct {
	side int
}

type rectangle struct {
	width  int
	length int
}

func (sq square) area() int {
	return sq.side * sq.side
}

func (rect rectangle) area() int {
	return rect.length * rect.width
}

func (sq square) perimeter() int {
	return 4 * sq.side
}

func (rect rectangle) perimeter() int {
	return 2 * (rect.length + rect.width)
}

func describe(i interface{}) {
	fmt.Printf("Type of String = %T, value of String = %v\n", i, i)
}

func main() {
	var rsh shape = rectangle{10, 20}
	var ssh shape = square{10}
	fmt.Println("Area of Square:", ssh.area())
	fmt.Println("Perimeter of Square: ", ssh.perimeter())
	fmt.Println("Area of Rectangle: ", rsh.area())
	fmt.Println("Perimeter of Rectangle: ", rsh.perimeter())
	s := "Empty Interface"
	describe(s)
	i := 100
	describe(i)
	strt := struct {
		name string
	}{
		name: "Sumathy",
	}
	describe(strt)
}

package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func MemoryOptimization() []string {
	Menus := []string{"Briyani", "Noodles", "Shawarma", "Icecream", "Brownie"}
	menuslice := Menus[:len(Menus)-3]
	fmt.Println("Menuslice(Sliced Array): ", menuslice)
	menucopy := make([]string, len(menuslice)) // copies the sliced array (menuslice) to menucopy
	copy(menucopy, menuslice)                  //copies the value of menuslice to menucopy
	return menucopy
}

func main() {
	cartoon := []string{"Brown", "Cony", "Peach", "Goma"}
	fmt.Println("cartoons:", cartoon, "has old length", len(cartoon), "and capacity", cap(cartoon))
	cartoon = append(cartoon, "Mocha")
	fmt.Println("cartoons:", cartoon, "has old length", len(cartoon), "and capacity", cap(cartoon))
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "Mathew", "Sumathy", "Sridhar")
		fmt.Println("names contents:", names)
	}
	food := []string{"Briyani", "Shawarma", "Icecream"}
	restaurants := []string{"Crescent", "Tasty Kitchen", "Pandias"}
	app := append(food, restaurants...)
	fmt.Println("Appended: ", app)
	nos := []int{1, 2, 3}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
	mdslice := [][]int{
		{10, 20},
		{30, 40},
		{50, 60}}
	fmt.Println("Multidimensional slice: ")
	for _, i := range mdslice {
		for _, j := range i {
			fmt.Printf("\t%d", j)
		}
		fmt.Printf("\n")
	}
	menus := MemoryOptimization()
	fmt.Println("Sliced Array : ", menus)
}

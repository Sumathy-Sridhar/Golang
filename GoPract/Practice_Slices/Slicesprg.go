package main

import (
	"fmt"
)

func main() {
	strarr := [5]string{"Brown", "Cony", "Delvin", "Milk", "Mocha"}
	fmt.Println("Original Array: ", strarr)
	var slicedarr []string = strarr[1:4] //creates a slice from a[1] to a[3]
	fmt.Println("Sliced Array:", slicedarr)
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println("Created Array:", c)
	darr := [...]int{1, 2, 3, 4, 5, 6}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after modifying:", darr)
	numarr := [3]int{1, 2, 3}
	numarr1 := numarr[:] //creates a slice which contains all elements of the array
	numarr2 := numarr[:]
	fmt.Println("array before change 1", numarr)
	numarr1[0] = 100
	fmt.Println("array after modification to slice numarr1", numarr)
	numarr2[1] = 1000
	fmt.Println("array after modification to slice numarr2", numarr)
	array := [...]string{"Sumi", "Sumathy", "Sridhar", "Mathew"}
	slice := array[1:3]
	fmt.Printf("Length of slice %d capacity %d\n", len(slice), cap(slice))
	slice = slice[:cap(slice)] //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(slice), "and capacity is", cap(slice))
	i := make([]int, 5, 5)
	fmt.Println(i)
}

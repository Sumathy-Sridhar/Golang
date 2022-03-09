package main

import "fmt"

func main() {

	//string array
	strarr := make([]string, 5)
	fmt.Println("Empty String array:", strarr)

	//assigning values to arr
	strarr[0] = "Brown"
	strarr[1] = "Cony"
	strarr[2] = "Peach"

	fmt.Println("String Array:", strarr)
	fmt.Println("Get:", strarr[1])
	fmt.Println("Len of array: ", len(strarr))

	strarr = append(strarr, "Goma")

	fmt.Println("String Array after appending:", strarr)

	//copy string array
	cpy := make([]string, len(strarr))
	copy(cpy, strarr)
	fmt.Println("Copy of array: ", cpy)

	//slicing
	sl := cpy[:2] // cpy[2:5] cpy[2:] cpy[0:4]
	fmt.Println(sl)

}

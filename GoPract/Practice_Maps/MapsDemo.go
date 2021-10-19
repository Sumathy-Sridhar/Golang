package main

import "fmt"

func main() {
	Bookdetails := make(map[int]string) // empty map is created
	fmt.Println("Before Addding Elemensts: ", Bookdetails)
	Bookdetails[123] = "Revolution2020" //adding elements to the map
	Bookdetails[143] = "Your dreams are mine now"
	fmt.Println("Elements of Map: ", Bookdetails)
	AccInfo := map[int]string{
		12345678: "Sumathy",
		25346879: "Sridhar",
		36897452: "Brown",
		78945614: "Cony",
		85239741: "Delvin",
	}
	fmt.Println("Account Details: ", AccInfo)
	fmt.Println("Iterating the Elements of AccInfo")
	for key, value := range AccInfo { //iterating elements of map using for loop and range
		fmt.Printf("AccInfo[%d] = %s\n", key, value)
	}
	fmt.Println("map before deletion", AccInfo)
	delete(AccInfo, 25346879) // deleting items from map using key value
	fmt.Println("map after deletion", AccInfo)
	// var bankdet map[string]int
	// bankdet["Delvin"] = 10000 // assignment to nil map {Runtime error} ==>Panic
	bookid := 143
	bookname := Bookdetails[bookid]
	fmt.Printf("Book Name of Book Id %d is %s\n", bookid, bookname) //Retrieving value for a key from a map
	search := 12345678
	value, ok := AccInfo[search]
	if ok == true {
		fmt.Println("Account Number of", search, "is", value)
		return
	}
	fmt.Println("Key Not Found")

}

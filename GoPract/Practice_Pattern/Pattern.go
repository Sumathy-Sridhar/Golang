package main

import "fmt"

func Pattern(row int){
	for i:=1;i<=row;i++{
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main()  {
	var row int 
	fmt.Println("Enter row count: ")
	fmt.Scanln(&row)
	Pattern(row)
}
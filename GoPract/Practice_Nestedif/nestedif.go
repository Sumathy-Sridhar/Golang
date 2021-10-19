package main

import "fmt"

func main()  {
fmt.Println("Enter a: ")
    var a,num,c int
    fmt.Scanln(&a) // reads input without the need of format specifiers
    fmt.Println("Enter num: ")
    fmt.Scanln(&num)
    fmt.Println("Enter c: ")
    fmt.Scanln(&c)
	if a>num && a>c{
		fmt.Printf("%d is greater than %d and %d \n",a,num,c)
	}else if num>c{
		fmt.Printf("%d is greater than %d and %d \n ",num,a,c)
	}else{
		fmt.Printf("%d is greater than %d and %d \n ",c,a,num)
	}

	if n:=10; n%2==0{ // assignment with if 
		fmt.Println(n, "is even")
	}else{
		fmt.Println(n,"is odd")
	}
}
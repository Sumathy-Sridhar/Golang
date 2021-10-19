package main

import "fmt"

func factorial(n int ) int {
	fact:=1
	for i:=1;i<=n;i++{
		fact=fact*i
	}
	return fact
}
func main()  {
	var n,res int
	fmt.Println("Enter n value : ")
	fmt.Scanln(&n)
	res=factorial(n)
	fmt.Printf("Factorial of %d is %d: \n",n,res)
	fmt.Println("Use of Break in For Loop")
	for j:=1;j<=10;j++{
		if j>5{
			break
		}
		fmt.Printf("%d\n",j)
	}
	fmt.Println("End of Loop...")
	fmt.Println("Use of Continue in For Loop")
	for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%d\n ", i)
    }
	fmt.Println("Without Initialization nd increment in forloop itself" )
	k:=0
	for k<=5{
		fmt.Printf("%d\n",k)
		k+=2
	}
	fmt.Println("Multiple Declarations in for loop")
	for no, i := 5, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
        fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
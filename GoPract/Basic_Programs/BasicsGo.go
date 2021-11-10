package main

import (
	"fmt"
	"strings"
)

func OddEven() {
	fmt.Println("Enter number to check odd or even: ")
	var num int
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("%d is even", num)
	} else {
		fmt.Printf("%d is odd", num)
	}
}
func LargestOfThree() {
	var (
		a int
		b int
		c int
	)
	fmt.Println("Enter a: ")
	fmt.Scanln(&a)
	fmt.Println("Enter b: ")
	fmt.Scanln(&b)
	fmt.Println("Enter c: ")
	fmt.Scanln(&c)

	if (a > b) && (a > c) {
		fmt.Printf("%d is greater than %d and %d", a, b, c)
	} else if b > c {
		fmt.Printf("%d is greater than %d and %d", b, a, c)
	} else {
		fmt.Printf("%d is greater than %d and %d", c, a, b)
	}
}

func MultiplicationTable() {
	var (
		number int
		i      int
		res    int
	)
	fmt.Println("Enter number to generate multiplication table: ")
	fmt.Scanln(&number)
	fmt.Printf("Multiplication table of %d\n", number)
	for i = 1; i <= 10; i++ {
		res = number * i
		fmt.Printf("%d*%d=%d\n", number, i, res)
	}

}
func Factorial() {
	var (
		i    int
		fnum int
	)
	fact := 1
	fmt.Println("Enter number to find factorial: ")
	fmt.Scanln(&fnum)
	if fnum < 0 {
		fmt.Println("Negative numbers not allowed")
	} else {
		for i = 1; i <= fnum; i++ {
			fact = fact * i
		}
		fmt.Printf("%d", fact)
	}
}
func ArmstrongNumber() {
	var (
		armnum    int
		tempo     int
		remainder int
		result    int = 0
	)
	fmt.Println("Enter number to check if its Armstrong number: ")
	fmt.Scanln(&armnum)
	tempo = armnum
	for {
		remainder = tempo % 10
		result += remainder * remainder * remainder
		tempo /= 10

		if tempo == 0 {
			break
		}
	}
	if result == armnum {
		fmt.Printf("%d is Armstrong Number", armnum)
	} else {
		fmt.Printf("%d is not an Armstrong number", armnum)
	}
}

func NaturalSum() {
	var (
		n   int
		sum int = 0
	)
	fmt.Println("Enter number : ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Sum of natural numbers: %d", sum)
}

func ReverseNumber() {
	var (
		pnum int
		rev  int = 0
		rem  int
		temp int
	)
	fmt.Println("Enter the number to reverse: ")
	fmt.Scanln(&pnum)
	temp = pnum
	for {
		rem = pnum % 10
		rev = (rev * 10) + rem
		pnum /= 10
		if pnum == 0 {
			break
		}
	}
	fmt.Printf("Reverse of %d is %d ", temp, rev)
}

func StringConcat() {
	var (
		str1 string
		str2 string
	)
	fmt.Println("Enter String1: ")
	fmt.Scanln(&str1)
	fmt.Println("Enter String2: ")
	fmt.Scanln(&str2)
	fmt.Println("Concated String : ", str1+str2)
}

func CircleAreaCircum() {
	var (
		radius        float32
		area          float32
		circumference float32
	)
	const pi = 3.14
	fmt.Println("Enter radius: ")
	fmt.Scanln(&radius)
	area = pi * radius * radius
	fmt.Printf("Area of circle with radius %f is: %f\n", radius, area)
	circumference = 2 * pi * radius
	fmt.Printf("Circumference of circle with radius %f is: %f", radius, circumference)
}
func FibonacciSeries() {
	var (
		n int
		a int = -1
		b int = 1
		c int = 0
		i int
	)
	fmt.Println("Enter the length of fibonacci series: ")
	fmt.Scanln(&n)
	for i = 1; i <= n; i++ {
		c = a + b
		a = b
		b = c
		fmt.Println(c)
	}
}

func StringFunctions() {
	// case sensitive
	fmt.Println("Contains Fn: ", strings.Contains("Sumathy", "mathy"))
	fmt.Println("Count Function: ", strings.Count("Malayalam", "A"))
	fmt.Println("Count Function: ", strings.Count("Malayalam", "a"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("Test Drive", "te"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("Test Drive", "Te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("Tesla", "la"))
	fmt.Println("Index Function: ", strings.Index("Tesla", "la"))
	fmt.Println("Index Function: ", strings.Index("Tesla", "a"))
	fmt.Println("Join Function: ", strings.Join([]string{"Sumathy", "Sridhar"}, "-"))
	fmt.Println("Repeat Function: ", strings.Repeat("Delvin  ", 3))
	fmt.Println("Replace Function: ", strings.Replace("Sumathi", "i", "y", 1))
	fmt.Println("Replace Function: ", strings.Replace("Sumuthi", "uthi", "athy", 1))
	fmt.Println("Split Function: ", strings.Split("Sumathy-Sridhar", "-"))
	fmt.Println("ToLower: ", strings.ToLower("SumATHy"))
	fmt.Println("ToUpper: ", strings.ToUpper("SumATHy"))
	fmt.Println("Length: ", len("Swift"))
	fmt.Println("Character ASCII value : ", "Brown"[4])
	fmt.Println("Compare function: ", strings.Compare("Sumathy", "sumathy"))
	fmt.Println("Compare function: ", strings.Compare("Sumathy", "Sumathy"))
	fmt.Println("Compare function: ", strings.Compare("sumathy", "Sumathy"))
	fmt.Println(strings.ContainsAny("Germany", "G"))
	fmt.Println(strings.ContainsAny("Germany", "g"))
	fmt.Println(strings.EqualFold("Cat", "cAt"))
	fmt.Println(strings.EqualFold("India", "Indiana"))
}
func MatrixMultiplication() {
	var (
		matrx1 [10][10]int
		matrx2 [10][10]int
		sum    [20][20]int
		row    int
		col    int
	)
	fmt.Println("Enter no of rows: ")
	fmt.Scanln(&row)
	fmt.Println("Enter no of column: ")
	fmt.Scanln(&col)
	fmt.Println("====MAtrix 1=====")
	for i := 0; i <= row; i++ {
		for j := 0; j <= col; j++ {
			fmt.Printf("Enter the element of matrix1 %d %d :", i+1, j+1)
			fmt.Scanln(&matrx1[i][j])
		}
	}
	fmt.Println("====MAtrix 2=====")
	for i := 0; i <= row; i++ {
		for j := 0; j <= col; j++ {
			fmt.Printf("Enter the element of matrix2 %d %d :", i+1, j+1)
			fmt.Scanln(&matrx2[i][j])
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = matrx1[i][j] + matrx2[i][j]
		}
	}
	fmt.Println("========== Sum of Matrix =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ", sum[i][j])
			if j == col-1 {
				fmt.Println("")
			}
		}
	}
}
func Swap() {
	var (
		first  int
		second int
	)
	fmt.Println("Enter First number: ")
	fmt.Scanln(&first)
	fmt.Println("Enter Second Number: ")
	fmt.Scanln(&second)
	fmt.Println("===Before Swapping===")
	fmt.Println("First number: ", first)
	fmt.Println("Second Number: ", second)
	first = first + second
	second = first - second
	first = first - second
	fmt.Println("===After Swapping===")
	fmt.Println("First number: ", first)
	fmt.Println("Second Number: ", second)
}

func main() {
	// OddEven()
	// LargestOfThree()
	//MultiplicationTable()
	//Factorial()
	//ArmstrongNumber()
	//NaturalSum()
	//ReverseNumber()
	//StringConcat()
	//CircleAreaCircum()
	//FibonacciSeries()
	//StringFunctions()
	//MatrixMultiplication()
	Swap()
}

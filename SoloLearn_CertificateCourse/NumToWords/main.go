package main

import "fmt"

func main() {

	var num int

	for i := 1; i <= 3; i++ {

		fmt.Scanln(&num)

		switch num {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		case 10:
			fmt.Println("Ten")
		default:
			fmt.Println("Number out of Range")
		}
	}
}

// a := 8
// b := 3
// fmt.Println(a % b)

// // x := 6
// // y := x - 1
// // x %= y
// // fmt.Println(x)

// // age := 15
// // res := (age > 18) || age == 0
// // fmt.Println(res)

// // if x := 8; x/2 != 0 {
// // 	fmt.Println(x - 2)
// // } else {
// // 	fmt.Println(x)
// // }

// // x := 8
// // switch y := x % 2; y {
// // case 0:
// // 	x -= 1
// // case 1:
// // 	x += 1
// // }
// // fmt.Println(x)

// // sum := 0
// // for i := 1; i <= 3; i++ {
// // 	sum += i
// // }
// // fmt.Println(sum)

// sum := 0
// for sum < 5 {
// 	sum += 2
// }
// fmt.Println(sum)

// // x := 7
// // res := 0
// // switch {
// // case x > 8:
// // 	res++
// // case x > 3 && x < 10:
// // 	res++
// // case x > 5:
// // 	res++
// // }
// // fmt.Println(res)

// res := 0
// for i := 0; i < 10; i += 3 {
// 	res += i
// }
// fmt.Println(res)

package main

import "fmt"

// // func convert(x int) {
// // 	fmt.Println(x / 100)
// // }
// // func main() {
// // 	convert(300)
// // 	defer fmt.Println(1)
// // 	fmt.Println(2)
// // }

// // var x = 3

// // func change() {
// // 	x += 1
// // }
// // func set(x int) {
// // 	x += 1
// // }
// // func main() {
// // 	change()
// // 	set(x)
// // 	fmt.Println(x)
// // }

// // func calc(a int) (int, int) {
// // 	return a + 2, a + 1
// // }
// // func main() {
// // 	x, y := calc(5)
// // 	fmt.Println(x + y)
// // }

// // func double(a int) {
// // 	fmt.Println(a * 2)
// // }
// // func main() {
// // 	for i := 4; i > 0; i-- {
// // 		defer double(i)
// // 	}
// // }
// // func main() {
// // 	a := 4
// // 	p := &a
// // 	a += 2
// // 	*p = *p - 1
// // 	fmt.Println(a)
// // }
// // type Coord struct {
// // 	x int
// // 	y int
// // }

// func main() {
// // 	a := Coord{7, 5}
// // 	fmt.Println(a.x - a.y)
// // }

// // type Test struct {
// // 	a int
// // 	b int
// // }

// // func (x Test) do() int {
// // 	return x.a - x.b
// // }
// // func main() {
// // 	t := Test{8, 3}
// // 	fmt.Println(t.do())
// // }

// type T struct {
// 	val int
// }

// func (p *T) a() {
// 	p.val += 1
// }
// func (p T) b() {
// 	p.val += 2
// }
// func main() {
// 	x := T{5}
// 	x.a()
// 	x.b()
// 	fmt.Println(x.val)
// }

// x := make([]int, 2)
// x = append(x, 3, 6, 2)
// fmt.Println(x[4])

// res := 0
// nums := [3]int{2, 4, 6}
// for i, v := range nums {
// 	if i%2 == 0 {
// 		res += v
// 	}
// }
// fmt.Println(res)

// m := map[int]int{
// 	8: 42,
// 	2: 6,
// 	4: 9,
// 	5: 3}
// delete(m, 2)
// fmt.Println(m[4] - m[5])

// func f(v ...int) {
// 	res := 20
// 	for _, a := range v {
// 		res -= a
// 	}
// 	fmt.Println(res)
// }
// func main() {
// 	v := []int{8, 5, 3}
// 	f(v...)
// }

// func main() {
// 	s := []int{1, 2, 4, 6, 8}
// 	s[2] = s[1]
// 	s[3] = s[2] + s[0]
// 	fmt.Println(s[4] % s[3])

// // }

// func f(ch chan int) {
// 	ch <- 8
// }
// func main() {
// 	ch := make(chan int)
// 	go f(ch)
// 	select {
// 	case <-ch:
// 		fmt.Println("a")
// 	default:
// 		fmt.Println("b")
// 	}
// }

func f(ch chan int) {
	ch <- 4
}
func main() {
	ch := make(chan int)
	go f(ch)
	fmt.Println(<-ch / 2)
}

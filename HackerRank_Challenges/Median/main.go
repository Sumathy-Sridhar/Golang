package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr [50]int
	var n, len int
	len = n
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr[0:len])

	if n%2 != 0 {
		rem := (len + 1) / 2
		median := arr[rem-1]
		fmt.Print(median)
	}

}

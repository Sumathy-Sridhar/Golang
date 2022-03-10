package main

import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	var sum int
	var val string

	fmt.Scanln(&val)
	results = append(results, val)

	for _, v := range results {
		switch {
		case v == "w":
			sum += 3
		case v == "l":
			sum += 0
		case v == "d":
			sum += 1
		}
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {

	var arrlen, pos, neg, zero, arrnum int

	fmt.Scan(&arrlen)

	for i := 1; i <= arrlen; i++ {
		fmt.Scan(&arrnum)
		if arrnum > 0 {
			pos++
		} else if arrnum < 0 {
			neg++
		} else {
			zero++
		}
	}
	posRatio := float64(pos) / float64(arrlen)
	negRatio := float64(neg) / float64(arrlen)
	zeroRatio := float64(zero) / float64(arrlen)

	fmt.Printf("%.6f\n", posRatio)
	fmt.Printf("%.6f\n", negRatio)
	fmt.Printf("%.6f\n", zeroRatio)

}

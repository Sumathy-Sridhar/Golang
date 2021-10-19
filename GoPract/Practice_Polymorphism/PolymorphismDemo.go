package main

import "fmt"

type TaxInterface interface {
	calculateTax() int
}

type IndianTax struct {
	taxPercentage int
	income        int
}

func (ind *IndianTax) calculateTax() int {
	tax := ind.income * ind.taxPercentage / 100
	return tax
}

type UsaTax struct {
	taxPercentage int
	income        int
}

func (us *UsaTax) calculateTax() int {
	tax := us.income * us.taxPercentage / 50
	return tax
}

func main() {
	indtax := &IndianTax{
		taxPercentage: 30,
		income:        1000,
	}
	res := indtax.calculateTax()
	fmt.Println("Indian Tax: ", res)
	ustax := &UsaTax{
		taxPercentage: 80,
		income:        2000,
	}
	res1 := ustax.calculateTax()
	fmt.Println("USA Tax: ", res1)
}

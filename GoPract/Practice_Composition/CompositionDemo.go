package main

import "fmt"

type Customer struct {
	name  string
	accno int
}

type Bank struct {
	balance float32
	limit   float32
	Customer
}

func (b Bank) bankDetails() {
	fmt.Println("****CUSTOMER DETAILS*****")
	fmt.Println("Customer Name: ", b.name)
	fmt.Println("Account Number: ", b.accno)
	fmt.Println("Balance:", b.balance)
	fmt.Println("Credit Limit:", b.limit)
}

type bankSlice struct {
	banks []Bank
}

func (bs bankSlice) DisplayAll() {
	fmt.Println("\n Printing Customer Details for past 3 years..")
	for _, v := range bs.banks {
		v.bankDetails()
		fmt.Println()
	}
}

func main() {
	cus := Customer{
		"Delvin",
		123456789,
	}
	bankvar1 := Bank{
		80000.00,
		20000.50,
		cus,
	}
	bankvar2 := Bank{
		90000.00,
		10000.75,
		cus,
	}
	bankvar3 := Bank{
		100000.00,
		50000.25,
		cus,
	}
	bs := bankSlice{
		banks: []Bank{bankvar1, bankvar2, bankvar3},
	}
	bs.DisplayAll()
}

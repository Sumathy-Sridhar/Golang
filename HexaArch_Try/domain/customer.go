package domain

type Customer struct {
	Customerid   int
	Customername string
	Address      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Customerid: 123, Customername: "Sumathy", Address: "mn nagar"},
		{Customerid: 365, Customername: "Delvin", Address: "hhfg street"},
	}
	return CustomerRepositoryStub{customers: customers}
}

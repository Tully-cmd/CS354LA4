package bank

type ICustomer interface {
	ToString() string
	NewCustomer(string) Customer
}

type Customer struct {
	name string
}

func NewCustomer(name string) (Customer) {
	var c Customer
	c.name = name
	return c
}

func (c Customer) ToString() string {
	return c.name
}




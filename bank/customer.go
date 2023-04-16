package bank

type ICustomer interface {
	ToString() string
	//NewCustomer(string) Customer
	Init(string) 
}

type Customer struct {
	name string
}

func NewCustomer(name string) (c *Customer) {
	c=new(Customer)
	c.Init(name)
	return
}

func (c *Customer) Init(name string) {
	c.name=name
}

func (c *Customer) ToString() string {
	return c.name
}




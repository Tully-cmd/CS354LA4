package bank

import . "fmt"
//import . "strconv"

type IAccount interface {
	Balance () float64
	Deposit (float64)
	Withdraw (float64)
	ToString () string
	Accure (float32)
}

type Account struct {
	customer Customer
	number string
	balance float64
}

// func NewCustomer(name string) (c *Customer) {
// 	c=new(Customer)
// 	c.Init(name)
// 	return
// }

// func (c *Customer) Init(name string) {
// 	c.name=name
// }

func (a Account) Balance () float64 {
	return a.balance
}

func (a Account) Deposit (amount float64) {
	a.balance += amount
}

func (a Account) Withdraw (amount float64) {
	a.balance -= amount
}

func (a Account) Accure (rate float32) {
	
}

func (a Account) ToString () string {

	return Sprintf("%s:%s:%.2f",a.number,a.customer.ToString(),a.balance)
	//return strconv.Itoa(a.number) + ":" + a.customer.ToString() + ":" + strconv.Itoa(a.balance)
}
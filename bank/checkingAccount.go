package bank

type ICheckingAccount interface {
	IAccount
	NewCheckingAccount(string,Customer,float64) (*CheckingAccount)
}

type CheckingAccount struct {
	Account
}

func NewCheckingAccount(number string, customer Customer, balance float64) (CheckingAccount) {
	var a CheckingAccount
	a.number = number
	a.customer = customer
	a.balance = balance
	//a.Init(number,customer,balance)
	return a
}
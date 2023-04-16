package bank

type ICheckingAccount interface {
	IAccount
}

type CheckingAccount struct {
	Account
}

func NewCheckingAccount(number string, customer Customer, balance float64) (a *CheckingAccount) {
	a=new(CheckingAccount)
	a.Init(number,customer,balance)
	return
}

func (a *CheckingAccount) Init(number string, customer Customer, balance float64) {
	a.number = number
	a.customer = customer
	a.balance = balance
}

func (a *CheckingAccount) ToString() string {
	return a.ToString()
}
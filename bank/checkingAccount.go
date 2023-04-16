package bank

type ICheckingAccount interface {
	IAccount
}

type CheckingAccount struct {
	Account
}

func NewAccount(string name, Customer customer, float64 balance) (a *CheckingAccount {
	a=new(CheckingAccount)
	a.Init(name,customer,balance);
	return
}

func (a *CheckingAccount) Init(string name, Customer customer, float64 balance) {
	a.Account.name = name
	a.Account.Customer = customer
	a.Account.balance = balance
}

func Accure(rate float32) {
	
}

func (a *CheckingAccount) ToString() string {
	return a.Account.ToString()
}
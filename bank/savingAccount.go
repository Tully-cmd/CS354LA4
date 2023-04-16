package bank

type ISavingAccount interface {
	IAccount
	interest float32
}

type SavingAccount struct {
	Account
}

func NewAccount(string name, Customer customer, float64 balance) (a *SavingAccount {
	a=new(SavingAccount)
	a.Init(name,customer,balance);
	return
}

func (a *SavingAccount) Init(string name, Customer customer, float64 balance) {
	a.Account.name = name
	a.Account.Customer = customer
	a.Account.balance = balance
}

func Accure(rate float32) {
	interest+=balance*rate
	balance+=balance*rate
}

func (a *SavingAccount) ToString() string {
	return a.Account.ToString()
}
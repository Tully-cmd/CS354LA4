package bank

type ISavingAccount interface {
	IAccount
}

type SavingAccount struct {
	Account
	interest float32
}

func NewSavingAccount(string name, Customer customer, float64 balance) (a *SavingAccount) {
	a=new(SavingAccount)
	a.Init(name,customer,balance);
	return
}

func (a *SavingAccount) Init(string name, Customer customer, float64 balance) {
	a.Account.name = name
	a.Account.Customer = customer
	a.Account.balance = balance
}

func AccureSaving(rate float32) {
	interest+=balance*rate
	balance+=balance*rate
}

func (a *SavingAccount) ToString() string {
	return a.Account.ToString()
}
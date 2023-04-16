package bank

type ISavingAccount interface {
	IAccount
	NewSavingAccount(string,Customer,float64)
	AccureSaving(float32)
}

type SavingAccount struct {
	Account
	interest float32
}

func NewSavingAccount(number string, customer Customer, balance float64) (a *CheckingAccount) {
	a=new(CheckingAccount)
	a.Init(number,customer,balance)
	return
}

func (a *SavingAccount) Init(number string, customer Customer, balance float64) {
	a.number = number
	a.customer = customer
	a.balance = balance
}

func (a *SavingAccount) AccureSaving(rate float32) {
	//a.interest += (a.balance * rate)
	//a.balance += (a.balance * rate)
}

func (a *SavingAccount) ToString() string {
	return a.ToString()
}

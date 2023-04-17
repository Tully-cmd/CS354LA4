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

func NewSavingAccount(number string, customer Customer, balance float64) (SavingAccount) {
	var a SavingAccount
	a.number = number
	a.customer = customer
	a.balance = balance
	return a
}

func (a SavingAccount) AccureSaving(rate float32) {
	//a.interest += (a.balance * rate)
	//a.balance += (a.balance * rate)
}

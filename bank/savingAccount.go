package bank

import . "fmt"

type ISavingAccount interface {
	IAccount
	NewSavingAccount(string,Customer,float64)
	Accure(float32)
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

func (a SavingAccount) Accure (rate float32) SavingAccount {
	Printf("Accuring SavingAccount %s with a rate of %.2f\n",a.ToString(),rate)
	a.interest = float32(float64(a.interest) + a.balance * float64(rate))
	a.balance = a.balance + a.balance * float64(rate)
	Printf("Balance After Accure: %.2f\n",a.balance)
	return a
}

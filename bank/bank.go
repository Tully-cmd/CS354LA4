package bank

import . "fmt"

type IBank interface {
	IAccount
	ToString() string
	Accure(float32)
	Add(Account)
	NewBank() Bank
}

type Bank struct {
	bank []Account 
}

func NewBank() (Bank) {
	var b Bank
	b.bank = nil
	return b
}

func (b Bank) Add(a Account) Bank {
	b.bank = append(b.bank,a)
	return b
}

func (b Bank) Accure(rate float32) Bank{
	for i, acc := range b.bank {
		i = i
		Printf("acc.balance: %.2f\n",acc.balance)
		b.bank[i] = acc.Accure(rate)
		Printf("acc.balance: %.2f\n",acc.balance)
	}
	Printf("b.ToString(): %s\n",b.ToString())
	return b
}

func (b Bank) ToString() string {
	var retString = ""

	for i, acc := range b.bank {
		i = i
		retString += acc.ToString()
		retString += "\n"
	}

	return retString
}
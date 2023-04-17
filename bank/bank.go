package bank

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

func (b Bank) Accure(rate float32) {
	for i, acc := range b.bank {
		i = i
		acc.Accure(rate)
	}
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
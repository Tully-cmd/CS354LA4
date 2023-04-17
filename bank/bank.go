package bank

type IBank interface {
	IAccount
	ToString() string
	Accure(float32)
	Add(Account)
	NewBank() Bank
}

type Bank struct {
	checkings []CheckingAccount
	savings []SavingAccount 
}

func NewBank() (Bank) {
	var b Bank
	b.checkings = nil
	b.savings = nil
	return b
}

func (b Bank) AddChecking(a CheckingAccount) Bank {
	b.checkings = append(b.checkings,a)
	return b
}

func (b Bank) AddSaving(a SavingAccount) Bank {
	b.savings = append(b.savings,a)
	return b
}

func (b Bank) Accure(rate float32) Bank{
	for i, acc := range b.savings {
		i = i
		b.savings[i] = acc.Accure(rate)
	}
	return b
}

func (b Bank) ToString() string {
	var retString = "\nChecking Accounts\n\n"

	for i, acc := range b.checkings {
		i = i
		retString += acc.ToString()
		retString += "\n"
	}

	retString += "\nSaving Accounts\n\n"

	for i, acc := range b.savings {
		i = i
		retString += acc.ToString()
		retString += "\n"
	}

	return retString
}
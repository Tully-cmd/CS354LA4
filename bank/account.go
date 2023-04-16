package bank

type IAccount interface {
	ICheckingAccount
	ISavingAccount
}

type Account struct {
	Customer
	number string
	balance float64

	Accure(rate float32)

	func Balance () float64 {
		return balance
	}
	
	func Deposit (amount float64) {
		balance += amount
	}

	func Withdraw (amount float64) {
		balance -= amount
	}

	func ToString () string {
		Printf("%d:%s:%d",number,Customer,balance)
	}
}
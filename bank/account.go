package bank

import . "fmt"

type IAccount interface {
	Accure(rate float32)
}

type Account struct {
	customer Customer
	number string
	balance float64
}

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
	Printf("%d:%s:%d",number,customer,balance)
	return number + ":" + customer.ToString() + ":" + balance
}
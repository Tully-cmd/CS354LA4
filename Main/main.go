package main

import . "fmt"
import . "local/bank"

func main() {
	var customer Customer=NewCustomer("Tully")
	var checkingAccount CheckingAccount=NewCheckingAccount("1111",customer,4000)
	var savingAccount SavingAccount=NewSavingAccount("1112",customer,800)

	Printf("customer.ToString() => %s\n",customer.ToString())
	Printf("checkingAccount.ToString() => %s\n",checkingAccount.ToString())
	Printf("savingAccount.ToString() => %s\n\n",savingAccount.ToString())

	Printf("Withdrawing 200 from Account #%s\n",checkingAccount.ToString())
	checkingAccount.Account = checkingAccount.Account.Withdraw(200)
	Printf("Depositing 50 into Account #%s\n",savingAccount.ToString())
	savingAccount.Account = savingAccount.Account.Deposit(50)

	var bank Bank=NewBank()
	bank = bank.Add(checkingAccount.Account)
	bank = bank.Add(savingAccount.Account)
	bank = bank.Accure(0.03)
	Printf("bank.ToString() => %s\n",bank.ToString())	
}
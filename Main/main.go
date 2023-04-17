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


	var bank Bank=NewBank()
	bank = bank.Add(checkingAccount.Account)
	bank = bank.Add(savingAccount.Account)
	bank.Accure(0.03)
	Printf("bank.ToString() => %s\n",bank.ToString())

}
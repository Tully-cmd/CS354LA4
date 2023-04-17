package main

import . "fmt"
import . "local/bank"

func main() {
	var customer Customer=NewCustomer("Tully")
	var checkingAccount CheckingAccount=NewCheckingAccount("1111",customer,4000)
	var savingAccount SavingAccount=NewSavingAccount("1112",customer,800)
	
	var mary Customer=NewCustomer("Mary")
	var maryChecking CheckingAccount=NewCheckingAccount("1113",mary,10000)
	var marySaving SavingAccount=NewSavingAccount("1114",mary,50000)

	var mariah Customer=NewCustomer("Mariah")
	var mariahChecking CheckingAccount=NewCheckingAccount("1115",mariah,400)
	var mariahSaving SavingAccount=NewSavingAccount("1116",mariah,30000)

	var ivy Customer=NewCustomer("Ivy")
	var ivyChecking CheckingAccount=NewCheckingAccount("1117",ivy,2000)
	var ivySaving SavingAccount=NewSavingAccount("1118",ivy,15000)
	
	var brynn Customer=NewCustomer("Brynn")
	var brynnChecking CheckingAccount=NewCheckingAccount("1119",brynn,8000)
	var brynnSaving SavingAccount=NewSavingAccount("1120",brynn,4000)

	Printf("\ncustomer.ToString() => %s\n",customer.ToString())
	Printf("checkingAccount.ToString() => %s\n",checkingAccount.ToString())
	Printf("savingAccount.ToString() => %s\n\n",savingAccount.ToString())

	Printf("Withdrawing 200 from Account #%s\n",checkingAccount.ToString())
	checkingAccount.Account = checkingAccount.Account.Withdraw(200)
	Printf("Account after withdrawing #%s\n",checkingAccount.ToString())
	Printf("Depositing 50 into Account #%s\n",savingAccount.ToString())
	savingAccount.Account = savingAccount.Account.Deposit(50)
	Printf("Account after depositing #%s\n\n",savingAccount.ToString())

	var bank Bank=NewBank()
	bank = bank.AddChecking(checkingAccount)
	bank = bank.AddSaving(savingAccount)
	bank = bank.AddChecking(maryChecking)
	bank = bank.AddSaving(marySaving)
	bank = bank.AddChecking(mariahChecking)
	bank = bank.AddSaving(mariahSaving)
	bank = bank.AddChecking(ivyChecking)
	bank = bank.AddSaving(ivySaving)
	bank = bank.AddChecking(brynnChecking)
	bank = bank.AddSaving(brynnSaving)
	bank = bank.Accure(0.03)
	Printf("\nbank.ToString() => %s\n",bank.ToString())	
}
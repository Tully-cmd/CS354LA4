package main

import . "fmt"
import . "local/bank"

func main() {
	var a Customer=NewCustomer("Tully")
	var aA CheckingAccount=NewCheckingAccount("1111",a,4000)

	Printf("a.ToString() => %s\n",a.ToString())
	Printf("aA.ToString() => %s\n",aA.ToString())
}
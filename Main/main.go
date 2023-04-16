package main

import . "fmt"
import . "local/bank"

func main() {
	var a ICustomer=NewCustomer("Tully")
	var aA ICheckingAccount=NewCheckingAccount("1111",a,4000)

	Printf("a.ToString() => %s\n",a.ToString())
	Printf("aA.ToString() => %s\n",aA.ToString())
	
}
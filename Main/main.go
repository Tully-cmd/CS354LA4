package main

import . "fmt"
import . "local/bank"

func main() {
	var a ICustomer=NewCustomer("Tully")
	//var aA ICheckingAccount=NewCheckingAccount("Tully",a,4000)

	Printf("a.ToString() => %s\n",a.ToString())
	//Printf("aA.ToString() => %s\n",aA.ToString())
	
}
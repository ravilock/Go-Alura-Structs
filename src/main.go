package main

import (
	"Go-Alura-Structs/src/account"
	"Go-Alura-Structs/src/bankServices"
	"fmt"
)

func main() {
	SomeOneAccount := account.CreateCheckingAccount("Some One", "10987654321", "Programmer", 20, 123, 12345)
	SomeOneAccount.Deposit(150)
	SomeOneAccount.Deposit(150)

	OneSomeAccount := account.CreateCheckingAccount("One Some", "12345678910", "Programmer", 20, 123, 12345)

	account.Transfer(SomeOneAccount, OneSomeAccount, 200)

	fmt.Println(SomeOneAccount, OneSomeAccount)

	bankServices.PayInvoice(OneSomeAccount, 200)

	fmt.Println(SomeOneAccount, OneSomeAccount)
}

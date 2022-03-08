package bankServices

import "Go-Alura-Structs/src/account"

func PayInvoice(account account.AccountInterface, invoiceValue float64) {
	account.Withdraw(invoiceValue)
}

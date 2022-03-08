package account

import "Go-Alura-Structs/src/accountHolder"

type SavingsAccount struct {
	accountHolder                                accountHolder.AccountHolder
	agencyNumber, accountNumber, operationNumber int
	balance                                      float64
}

func CreateSavingsAccount(name string, taxIdentifier string, profession string, age int, agencyNumber int, accountNumber int, operationNumber int) *SavingsAccount {
	newAccount := new(SavingsAccount)
	accountHolder := accountHolder.CreateAccountHolder(name, taxIdentifier, profession, age)
	newAccount.accountHolder = accountHolder
	newAccount.agencyNumber = agencyNumber
	newAccount.accountNumber = accountNumber
	newAccount.operationNumber = operationNumber
	return newAccount
}

func (account *SavingsAccount) Withdraw(withdrawValue float64) (float64, bool) {
	canWithdraw := withdrawValue <= account.balance && withdrawValue > 0
	if canWithdraw {
		account.balance -= withdrawValue
		return account.balance, true
	}
	return account.balance, false
}

func (account *SavingsAccount) Deposit(depositValue float64) (float64, bool) {
	canDeposit := depositValue > 0
	if canDeposit {
		account.balance += depositValue
		return account.balance, true
	}
	return account.balance, false
}

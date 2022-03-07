package account

import "Go-Alura-Structs/src/accountHolder"

type CheckingAccount struct {
	accountHolder               accountHolder.AccountHolder
	agencyNumber, accountNumber int
	balance                     float64
}

func CreateCheckingAccount(name string, taxIdentifier string, profession string, age int, agencyNumber int, accountNumber int) *CheckingAccount {
	newAccount := new(CheckingAccount)
	accountHolder := accountHolder.CreateAccountHolder(name, taxIdentifier, profession, age)
	newAccount.accountHolder = accountHolder
	newAccount.agencyNumber = agencyNumber
	newAccount.accountNumber = accountNumber
	return newAccount
}

func (account *CheckingAccount) Withdraw(withdrawValue float64) (float64, bool) {
	canWithdraw := withdrawValue <= account.balance && withdrawValue > 0
	if canWithdraw {
		account.balance -= withdrawValue
		return account.balance, true
	}
	return account.balance, false
}

func (account *CheckingAccount) Deposit(depositValue float64) (float64, bool) {
	canDeposit := depositValue > 0
	if canDeposit {
		account.balance += depositValue
		return account.balance, true
	}
	return account.balance, false
}

func Transfer(transferenceEmitter *CheckingAccount, transferenceReceiver *CheckingAccount, transferenceValue float64) bool {
	_, withdrawResult := transferenceEmitter.Withdraw(transferenceValue)
	if withdrawResult {
		_, depositResult := transferenceReceiver.Deposit(transferenceValue)
		return depositResult
	}
	return false
}

func (account *CheckingAccount) GetAccountBalance() float64 {
	return account.balance
}

package account

type AccountInterface interface {
	Withdraw(withdrawValue float64) (float64, bool)
	Deposit(depositValue float64) (float64, bool)
}

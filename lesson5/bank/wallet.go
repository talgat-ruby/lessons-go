package bank

type wallet struct {
	balance int
}

func (w *wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *wallet) Withdraw(amount int) {
	w.balance -= amount
}

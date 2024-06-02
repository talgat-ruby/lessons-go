package bank

type Account struct {
	name   string
	wallet *wallet
}

func NewAccount(name string) *Account {
	wallet := &wallet{0}
	return &Account{name: name, wallet: wallet}
}

func (a *Account) Deposit(amount int) {
	a.wallet.Deposit(amount)
}

func (a *Account) Withdraw(amount int) {
	newBalance := a.wallet.balance - amount

	if newBalance >= 0 {
		a.wallet.Withdraw(amount)
	}
}

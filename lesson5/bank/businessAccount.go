package bank

type BusinessAccount struct {
	name   string
	wallet *wallet
}

const MIN_NEGATIVE_BALANCE = -10000

func NewBusinessAccount(name string) *Account {
	wallet := &wallet{0}
	return &Account{name: name, wallet: wallet}
}

func (b *BusinessAccount) Deposit(amount int) {
	b.wallet.Deposit(amount)
}

func (b *BusinessAccount) Withdraw(amount int) {
	newBalance := b.wallet.balance - amount

	if newBalance >= -MIN_NEGATIVE_BALANCE {
		b.wallet.Withdraw(amount)
	}
}

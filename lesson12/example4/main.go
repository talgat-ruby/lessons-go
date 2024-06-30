package main

import (
	"fmt"
	"math/rand"
	"time"
)

type budget struct {
	balance int
	r       *rand.Rand
}

func newBudget() *budget {
	return &budget{
		balance: 1000,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (b *budget) deposit(amount int) {
	time.Sleep(time.Duration(b.r.Intn(500)+500) * time.Millisecond)
	b.balance += amount
}

func (b *budget) withdraw(amount int) {
	time.Sleep(time.Duration(b.r.Intn(500)+500) * time.Millisecond)
	b.balance -= amount
}

func (b *budget) getBalance() int {
	time.Sleep(time.Duration(b.r.Intn(500)+500) * time.Millisecond)
	return b.balance
}

func main() {
	bdg := newBudget()

	go bdg.deposit(100)
	go bdg.withdraw(200)
	go bdg.deposit(300)
	go bdg.withdraw(400)
	go bdg.withdraw(2000)
	go bdg.deposit(100)

	fmt.Println("balance:", bdg.getBalance())
}

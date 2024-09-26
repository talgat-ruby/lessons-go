package population

import (
	"fmt"
)

type Population struct {
	Amount uint
}

func (p *Population) info() {
	fmt.Printf("amount: %d\n", p.Amount)
}

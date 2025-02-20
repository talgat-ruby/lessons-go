package request

type CtrlCreateExpenseRequest struct {
	amount   int64
	category string
}

func NewCtrlCreateExpenseRequest(amount int64, category string) *CtrlCreateExpenseRequest {
	return &CtrlCreateExpenseRequest{
		amount:   amount,
		category: category,
	}
}

func (r *CtrlCreateExpenseRequest) GetAmount() int64 {
	return r.amount
}

func (r *CtrlCreateExpenseRequest) GetCategory() string {
	return r.category
}

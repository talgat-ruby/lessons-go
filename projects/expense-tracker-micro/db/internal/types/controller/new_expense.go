package controller

type NewExpenseReq interface {
	GetAmount() int64
	GetCategory() string
}

type NewExpenseResp interface{}

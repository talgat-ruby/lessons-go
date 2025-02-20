package controller

type AlterExpenseReq interface {
	GetID() string
	GetAmount() *int64
	GetCategory() *string
}

type AlterExpenseResp interface{}

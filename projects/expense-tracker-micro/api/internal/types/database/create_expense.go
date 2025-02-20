package database

type CreateExpenseReq interface {
	GetUserID() string
	GetAmount() int64
	GetCategory() string
}

type CreateExpenseResp interface{}

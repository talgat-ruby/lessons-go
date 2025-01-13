package database

type UpdateExpenseReq interface {
	GetID() string
	GetAmount() *int64
	GetCategory() *string
}

type UpdateExpenseResp interface{}

package database

type UpdateExpenseReq interface {
	GetUserID() string
	GetID() string
	GetAmount() *int64
	GetCategory() *string
}

type UpdateExpenseResp interface{}

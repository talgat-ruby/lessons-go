package database

type DeleteExpenseReq interface {
	GetUserID() string
	GetID() string
}

type DeleteExpenseResp interface{}

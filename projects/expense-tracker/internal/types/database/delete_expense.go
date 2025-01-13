package database

type DeleteExpenseReq interface {
	GetID() string
}

type DeleteExpenseResp interface{}

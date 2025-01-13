package database

import "context"

type User interface {
	CreateUser(context.Context, CreateUserReq) (CreateUserResp, error)
	FindUserByEmail(context.Context, FindUserByEmailReq) (FindUserByEmailResp, error)
}

type Expense interface {
	CreateExpense(context.Context, CreateExpenseReq) (CreateExpenseResp, error)
	UpdateExpense(context.Context, UpdateExpenseReq) (UpdateExpenseResp, error)
	DeleteExpense(context.Context, DeleteExpenseReq) (DeleteExpenseResp, error)
}

type Database interface {
	User
	Expense
}

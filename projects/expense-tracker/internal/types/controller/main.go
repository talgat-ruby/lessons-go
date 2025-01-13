package controller

import "context"

type Interceptor interface {
	Authenticator(context.Context, string) (context.Context, error)
}

type Auth interface {
	SignUp(context.Context, SignUpReq) (SignUpResp, error)
	SignIn(context.Context, SignInReq) (SignInResp, error)
}

type Expense interface {
	// ListExpense(context.Context, ListExpenseReq) (ListExpenseResp, error)
	NewExpense(context.Context, NewExpenseReq) (NewExpenseResp, error)
	AlterExpense(context.Context, AlterExpenseReq) (AlterExpenseResp, error)
	RemoveExpense(context.Context, RemoveExpenseReq) (RemoveExpenseResp, error)
}

type Controller interface {
	Interceptor
	Auth
	Expense
}

package database

import "context"

type User interface {
	CreateUser(context.Context, CreateUserReq) (CreateUserResp, error)
	FindUserByEmail(context.Context, FindUserByEmailReq) (FindUserByEmailResp, error)
}

type Database interface {
	User
}

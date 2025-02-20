package database

import "time"

type CreateUserReq interface {
	GetEmail() string
	GetPasswordHash() string
	GetSalt() string
}

type CreateUserResp interface {
	GetID() string
	GetEmail() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

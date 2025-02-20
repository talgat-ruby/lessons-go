package database

import "time"

type FindUserByEmailReq interface {
	GetEmail() string
}

type FindUserByEmailResp interface {
	GetID() string
	GetEmail() string
	GetPasswordHash() string
	GetSalt() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

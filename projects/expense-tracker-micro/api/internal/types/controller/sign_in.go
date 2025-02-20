package controller

type SignInReq interface {
	GetEmail() string
	GetPassword() string
}

type SignInResp interface {
	GetToken() string
}

package controller

type SignUpReq interface {
	GetEmail() string
	GetPassword() string
}

type SignUpResp interface {
	GetToken() string
}

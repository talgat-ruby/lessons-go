package controller

import "context"

type Controller interface {
	// auth
	SignUp(context.Context, SignUpReq) (SignUpResp, error)
	SignIn(context.Context, SignInReq) (SignInResp, error)
}

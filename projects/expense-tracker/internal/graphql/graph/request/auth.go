package request

type CtrlSignInRequest struct {
	email string
	pass  string
}

func NewCtrlSignInRequest(email string, pass string) *CtrlSignInRequest {
	return &CtrlSignInRequest{
		email: email,
		pass:  pass,
	}
}

func (r *CtrlSignInRequest) GetEmail() string {
	return r.email
}

func (r *CtrlSignInRequest) GetPassword() string {
	return r.pass
}

type CtrlSignUpRequest struct {
	email string
	pass  string
}

func NewCtrlSignUpRequest(email string, pass string) *CtrlSignUpRequest {
	return &CtrlSignUpRequest{
		email: email,
		pass:  pass,
	}
}

func (r *CtrlSignUpRequest) GetEmail() string {
	return r.email
}

func (r *CtrlSignUpRequest) GetPassword() string {
	return r.pass
}

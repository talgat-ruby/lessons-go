package statusError

type StatusError struct {
	status int
	msg    string
}

func New(status int, msg string) error {
	return &StatusError{status, msg}
}

func (st *StatusError) Error() string {
	return st.msg
}

func (st *StatusError) Status() int {
	return st.status
}

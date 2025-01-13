package permission

type Error struct {
	mes string
}

func NewError(m string) error {
	return &Error{
		mes: m,
	}
}

func (e *Error) Error() string {
	return e.mes
}

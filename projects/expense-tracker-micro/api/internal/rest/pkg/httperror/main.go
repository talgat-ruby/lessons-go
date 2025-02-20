package httperror

import (
	"errors"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/permission"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/validation"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/pkg/ptr"
)

type Message struct {
	basic          string
	validation     string
	authentication string
	permission     string
}

func NewMessage(basic, validation, authentication, permission string) *Message {
	return &Message{
		basic:          basic,
		validation:     validation,
		authentication: authentication,
		permission:     permission,
	}
}

func (m *Message) HandleError(w http.ResponseWriter, err error) {
	var mes string
	var status int
	switch {
	case errors.As(err, ptr.ToPtr(new(validation.Error))):
		mes = m.validation
		status = http.StatusBadRequest
	case errors.As(err, ptr.ToPtr(new(authentication.Error))):
		mes = m.authentication
		status = http.StatusUnauthorized
	case errors.As(err, ptr.ToPtr(new(permission.Error))):
		mes = m.permission
		status = http.StatusForbidden
	default:
		mes = m.basic
		status = http.StatusInternalServerError
	}

	if mes == "" {
		mes = http.StatusText(status)
	}

	http.Error(w, mes, status)
}

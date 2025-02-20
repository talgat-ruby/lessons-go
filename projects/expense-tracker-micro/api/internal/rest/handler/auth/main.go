package auth

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
)

type Auth struct {
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(logger *slog.Logger, ctrl controller.Controller) *Auth {
	return &Auth{
		logger: logger,
		ctrl:   ctrl,
	}
}

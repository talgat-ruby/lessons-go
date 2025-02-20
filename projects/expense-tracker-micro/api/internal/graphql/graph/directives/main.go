package directives

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
)

type Directives struct {
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(logger *slog.Logger, ctrl controller.Controller) *Directives {
	return &Directives{
		logger: logger.With(slog.String("module", "directives")),
		ctrl:   ctrl,
	}
}

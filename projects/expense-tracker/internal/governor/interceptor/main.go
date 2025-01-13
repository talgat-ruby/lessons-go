package interceptor

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
)

type Interceptor struct {
	logger *slog.Logger
	conf   *config.Config
}

func New(conf *config.Config, logger *slog.Logger) *Interceptor {
	return &Interceptor{
		conf:   conf,
		logger: logger,
	}
}

package graph

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	conf   *config.APIGraphQLConfig
	ctrl   controller.Controller
	logger *slog.Logger
}

func NewResolver(conf *config.APIGraphQLConfig, ctrl controller.Controller, logger *slog.Logger) *Resolver {
	return &Resolver{
		conf:   conf,
		ctrl:   ctrl,
		logger: logger.With(slog.String("module", "resolver")),
	}
}

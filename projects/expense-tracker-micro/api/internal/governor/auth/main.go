package auth

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/database"
)

type Auth struct {
	conf   *config.Config
	logger *slog.Logger
	db     database.Database
}

func New(conf *config.Config, logger *slog.Logger, db database.Database) *Auth {
	return &Auth{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}

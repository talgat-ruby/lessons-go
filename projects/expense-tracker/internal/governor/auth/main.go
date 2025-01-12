package auth

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

type Auth struct {
	conf   *config.Config
	logger *slog.Logger
	db     database.Database
	val    *validator.Validate
}

func New(conf *config.Config, logger *slog.Logger, db database.Database, val *validator.Validate) *Auth {
	return &Auth{
		conf:   conf,
		logger: logger,
		db:     db,
		val:    val,
	}
}

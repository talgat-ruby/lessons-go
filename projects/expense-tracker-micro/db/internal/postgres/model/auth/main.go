package auth

import (
	"database/sql"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/config"
)

type Auth struct {
	conf   *config.PostgresConfig
	logger *slog.Logger
	db     *sql.DB
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Auth {
	return &Auth{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}

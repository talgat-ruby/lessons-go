package model

import (
	"database/sql"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/model/auth"
)

type Model struct {
	*auth.Auth
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Model {
	return &Model{
		Auth: auth.New(conf, logger.With(slog.String("component", "auth")), db),
	}
}

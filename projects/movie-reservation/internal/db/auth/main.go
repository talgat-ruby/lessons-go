package auth

import (
	"database/sql"
	"log/slog"
)

type Auth struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(db *sql.DB, logger *slog.Logger) *Auth {
	return &Auth{
		logger: logger,
		db:     db,
	}
}

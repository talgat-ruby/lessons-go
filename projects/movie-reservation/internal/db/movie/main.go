package movie

import (
	"database/sql"
	"log/slog"
)

type Movie struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(db *sql.DB, logger *slog.Logger) *Movie {
	return &Movie{
		logger: logger,
		db:     db,
	}
}

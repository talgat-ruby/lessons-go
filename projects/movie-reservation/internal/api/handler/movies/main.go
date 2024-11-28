package movies

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db"
)

type Movies struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Movies {
	return &Movies{
		logger: logger,
		db:     db,
	}
}

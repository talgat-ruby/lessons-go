package handler

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/api/handler/movies"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db"
)

type Handler struct {
	*movies.Movies
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Movies: movies.New(logger, db),
	}
}

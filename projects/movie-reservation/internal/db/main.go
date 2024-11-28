package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db/movie"
)

type DB struct {
	logger *slog.Logger
	pg     *sql.DB
	*movie.Movie
}

func New(logger *slog.Logger) (*DB, error) {
	pgsql, err := newPgSQL()
	if err != nil {
		return nil, err
	}

	return &DB{
		logger: logger,
		pg:     pgsql,
		Movie:  movie.New(pgsql, logger),
	}, nil
}

func newPgSQL() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

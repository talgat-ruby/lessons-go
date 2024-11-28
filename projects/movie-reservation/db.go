package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	logger *slog.Logger
	pg     *sql.DB
}

func newDB(logger *slog.Logger) (*DB, error) {
	pgsql, err := NewPgSQL()
	if err != nil {
		return nil, err
	}

	return &DB{
		logger: logger,
		pg:     pgsql,
	}, nil
}

func NewPgSQL() (*sql.DB, error) {
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

func (db *DB) Init(ctx context.Context) error {
	log := db.logger.With("method", "Init")

	stmt := `
CREATE TABLE IF NOT EXISTS movie (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    posterUrl TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "fail create table movie", "error", err)
		return err
	}

	log.InfoContext(ctx, "success create table movie")
	return nil
}

type ModelMovie struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PosterURL   string     `json:"poster_url"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (db *DB) FindMovies(ctx context.Context) ([]ModelMovie, error) {
	log := db.logger.With("method", "FindMovies")

	movies := make([]ModelMovie, 0)

	stmt := `
SELECT id, title, description, posterUrl, created_at, updated_at 
FROM movie
`

	rows, err := db.pg.QueryContext(ctx, stmt)
	if err != nil {
		log.ErrorContext(ctx, "fail to query table movie", "error", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		movie := ModelMovie{}

		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.PosterURL,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		); err != nil {
			log.ErrorContext(ctx, "fail to scan movie", "error", err)
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success create table movie")
	return movies, nil
}

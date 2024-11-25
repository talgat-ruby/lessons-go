package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"

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

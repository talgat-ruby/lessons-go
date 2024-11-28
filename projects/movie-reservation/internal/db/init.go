package db

import (
	"context"
)

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

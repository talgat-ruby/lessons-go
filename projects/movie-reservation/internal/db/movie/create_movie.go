package movie

import (
	"context"
	"database/sql"
	"errors"
)

func (m *Movie) CreateMovie(ctx context.Context, insertData *ModelMovie) (*ModelMovie, error) {
	log := m.logger.With("method", "CreateMovie")

	stmt := `
INSERT INTO movie (title, description, posterUrl)
VALUES ($1, $2, $3)
RETURNING id, title, description, posterUrl, created_at, updated_at
`

	row := m.db.QueryRowContext(ctx, stmt, insertData.Title, insertData.Description, insertData.PosterURL)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to insert to table movie", "error", err)
		return nil, err
	}

	movie := ModelMovie{}

	if err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.PosterURL,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.ErrorContext(ctx, "no values was found", "error", err)
			return nil, nil
		}
		log.ErrorContext(ctx, "fail to scan movie", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert to table movie")
	return &movie, nil
}

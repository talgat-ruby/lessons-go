package movie

import (
	"context"
)

func (m *Movie) FindMovie(ctx context.Context, id int64) (*ModelMovie, error) {
	log := m.logger.With("method", "FindMovie")

	stmt := `
SELECT id, title, description, posterUrl, created_at, updated_at 
FROM movie
WHERE id = $1
`

	row := m.db.QueryRowContext(ctx, stmt, id)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table movie", "error", err)
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
		log.ErrorContext(ctx, "fail to scan movie", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table movie")
	return &movie, nil
}

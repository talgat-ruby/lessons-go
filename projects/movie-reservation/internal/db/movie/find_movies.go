package movie

import (
	"context"
)

func (m *Movie) FindMovies(ctx context.Context) ([]ModelMovie, error) {
	log := m.logger.With("method", "FindMovies")

	movies := make([]ModelMovie, 0)

	stmt := `
SELECT id, title, description, posterUrl, created_at, updated_at 
FROM movie
`

	rows, err := m.db.QueryContext(ctx, stmt)
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

	log.InfoContext(ctx, "success query table movie")
	return movies, nil
}

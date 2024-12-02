package movie

import (
	"context"
)

func (m *Movie) UpdateMovie(ctx context.Context, id int64, insertData *ModelMovie) error {
	log := m.logger.With("method", "UpdateMovie", "id", id)

	stmt := `
UPDATE movie
SET title = $2, description = $3, posterUrl = $4
WHERE id = $1
`

	_, err := m.db.ExecContext(ctx, stmt, id, insertData.Title, insertData.Description, insertData.PosterURL)
	if err != nil {
		log.ErrorContext(ctx, "fail to update the table movie", "error", err)
		return err
	}

	log.InfoContext(ctx, "success update the table movie")
	return nil
}

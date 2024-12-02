package movie

import (
	"context"
)

func (m *Movie) DeleteMovie(ctx context.Context, id int64) error {
	log := m.logger.With("method", "DeleteMovie", "id", id)

	stmt := `
DELETE FROM movie
WHERE id = $1
`

	_, err := m.db.ExecContext(ctx, stmt, id)
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table movie", "error", err)
		return err
	}

	log.InfoContext(ctx, "success delete from the table movie")
	return nil
}

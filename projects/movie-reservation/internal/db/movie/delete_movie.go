package movie

import (
	"context"
	"fmt"
)

func (m *Movie) DeleteMovie(ctx context.Context, id int64) error {
	log := m.logger.With("method", "DeleteMovie", "id", id)

	stmt := `
DELETE FROM movie
WHERE id = $1
`

	res, err := m.db.ExecContext(ctx, stmt, id)
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table movie", "error", err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table movie", "error", err)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "movie with id was not found", "id", id)
		return fmt.Errorf("movie with id was not found")
	}

	log.InfoContext(ctx, "success delete from the table movie")
	return nil
}

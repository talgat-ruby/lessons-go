package auth

import (
	"context"
	"database/sql"
	"errors"
)

type AccessTokenInput struct {
	UserID string
}

func (m *Auth) AccessToken(ctx context.Context, inp *AccessTokenInput) (*ModelUser, error) {
	log := m.logger.With("method", "AccessToken")

	stmt := `
SELECT id, email
FROM user_
WHERE id = $1;
`

	row := m.db.QueryRowContext(ctx, stmt, inp.UserID)
	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table user", "error", err)
		return nil, err
	}

	user := ModelUser{}

	if err := row.Scan(
		&user.ID,
		&user.Email,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.ErrorContext(ctx, "no user found", "error", err)
			return nil, nil
		}

		log.ErrorContext(ctx, "fail to scan user", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success select user")
	return &user, nil
}

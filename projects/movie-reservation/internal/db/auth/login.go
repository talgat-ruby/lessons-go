package auth

import (
	"context"
	"database/sql"
	"errors"
)

type LoginInput struct {
	User *ModelUser
}

func (m *Auth) Login(ctx context.Context, email string) (*ModelUser, error) {
	log := m.logger.With("method", "Login")

	stmt := `
SELECT id, email, password_hash, salt
FROM user_
WHERE email = $1;
`

	row := m.db.QueryRowContext(ctx, stmt, email)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table user", "error", err)
		return nil, err
	}

	user := ModelUser{}

	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "no user found", "error", err)
			return nil, nil
		}

		log.ErrorContext(ctx, "fail to scan user", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success login user")
	return &user, nil
}

package auth

import (
	"context"
)

type RegisterInput struct {
	User *ModelUser
}

func (m *Auth) Register(ctx context.Context, inp *RegisterInput) (*ModelUser, error) {
	log := m.logger.With("method", "Register")

	stmt := `
INSERT INTO user_ (email, password_hash, salt)
VALUES ($1, $2, $3)
RETURNING id, email, password_hash, salt;
`

	row := m.db.QueryRowContext(ctx, stmt, inp.User.Email, inp.User.PasswordHash, inp.User.Salt)
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
		log.ErrorContext(ctx, "fail to scan user", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success register user")
	return &user, nil
}

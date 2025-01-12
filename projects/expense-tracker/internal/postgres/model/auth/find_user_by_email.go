package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/generated/expenses/public/model"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

func (m *Auth) FindUserByEmail(ctx context.Context, req database.FindUserByEmailReq) (database.FindUserByEmailResp, error) {
	log := m.logger.With(slog.String("handler", "FindUserByEmail"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	mdl := model.User{
		Email: req.GetEmail(),
	}
	stmt := table.User.
		SELECT(table.User.AllColumns).
		WHERE(table.User.Email.EQ(pg.String(req.GetEmail())))
	if err := stmt.QueryContext(ctx, m.db, &mdl); err != nil {
		switch {
		case errors.Is(err, qrm.ErrNoRows):
			log.WarnContext(ctx, "no rows found")
			return nil, nil
		default:
			log.ErrorContext(ctx, "failed to find user by email", slog.Any("error", err))
			return nil, fmt.Errorf("failed to find user by email: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
		slog.String("email", mdl.Email),
	)
	return &findUserByEmailResp{mdl}, nil
}

type findUserByEmailResp struct {
	model.User
}

func (resp *findUserByEmailResp) GetID() string {
	return fmt.Sprint(resp.ID)
}

func (resp *findUserByEmailResp) GetEmail() string {
	return resp.Email
}

func (resp *findUserByEmailResp) GetPasswordHash() string {
	return resp.PasswordHash
}

func (resp *findUserByEmailResp) GetSalt() string {
	return resp.Salt
}

func (resp *findUserByEmailResp) GetCreatedAt() time.Time {
	return resp.CreatedAt
}

func (resp *findUserByEmailResp) GetUpdatedAt() time.Time {
	return resp.UpdatedAt
}

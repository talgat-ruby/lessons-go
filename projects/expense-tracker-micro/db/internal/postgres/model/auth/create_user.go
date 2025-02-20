package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/postgres/generated/expenses/public/model"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/types/database"
)

func (m *Auth) CreateUser(ctx context.Context, req database.CreateUserReq) (database.CreateUserResp, error) {
	log := m.logger.With(slog.String("handler", "CreateUser"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	mdl := model.User{
		Email:        req.GetEmail(),
		PasswordHash: req.GetPasswordHash(),
		Salt:         req.GetSalt(),
	}
	stmt := table.User.
		INSERT(table.User.Email, table.User.PasswordHash, table.User.Salt).
		MODEL(mdl).
		RETURNING(table.User.AllColumns)
	if err := stmt.QueryContext(ctx, m.db, &mdl); err != nil {
		switch {
		case errors.Is(err, qrm.ErrNoRows):
			log.WarnContext(ctx, "no rows found")
			return nil, nil
		default:
			log.ErrorContext(ctx, "failed to insert user", slog.Any("error", err))
			return nil, fmt.Errorf("failed to insert user: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
		slog.String("email", mdl.Email),
	)
	return &createUserResp{mdl}, nil
}

type createUserResp struct {
	model.User
}

func (resp *createUserResp) GetID() string {
	return fmt.Sprint(resp.ID)
}

func (resp *createUserResp) GetEmail() string {
	return resp.Email
}

func (resp *createUserResp) GetCreatedAt() time.Time {
	return resp.CreatedAt
}

func (resp *createUserResp) GetUpdatedAt() time.Time {
	return resp.UpdatedAt
}

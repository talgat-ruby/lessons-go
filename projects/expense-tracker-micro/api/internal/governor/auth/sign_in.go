package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/validation"
)

func (r *Auth) SignIn(ctx context.Context, req controller.SignInReq) (controller.SignInResp, error) {
	log := r.logger.With(slog.String("handler", "SignIn"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	dbReq := newFindUserDBReq(req.GetEmail())
	dbResp, err := r.db.FindUserByEmail(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "find user db failed", slog.Any("error", err))
		return nil, fmt.Errorf("find user db failed %w", err)
	}
	if dbResp == nil {
		log.WarnContext(ctx, "user not found")
		return nil, validation.NewError("invalid credentials")
	}

	isValid, err := authentication.VerifyPassword(
		req.GetPassword(),
		r.conf.API.Pepper,
		dbResp.GetPasswordHash(),
		dbResp.GetSalt(),
	)
	if err != nil {
		log.ErrorContext(ctx, "verify password failed", slog.Any("error", err))
		return nil, fmt.Errorf("verify password failed %w", err)
	}
	if !isValid {
		log.WarnContext(ctx, "invalid password")
		return nil, validation.NewError("invalid credentials")
	}

	u := &authentication.UserData{
		ID:    dbResp.GetID(),
		Email: dbResp.GetEmail(),
	}
	t, err := authentication.GenerateToken(u, r.conf.API.TokenSecret)
	if err != nil {
		log.ErrorContext(ctx, "generate token failed", slog.Any("error", err))
		return nil, fmt.Errorf("generate token failed %w", err)
	}

	log.InfoContext(
		ctx,
		"success",
		slog.String("email", dbResp.GetEmail()),
	)
	return newSignInResp(t), nil
}

type findUserDBReq struct {
	email string
}

func newFindUserDBReq(email string) *findUserDBReq {
	return &findUserDBReq{
		email: email,
	}
}

func (req *findUserDBReq) GetEmail() string {
	return req.email
}

type signInResp struct {
	token string
}

func newSignInResp(token string) *signInResp {
	return &signInResp{token: token}
}

func (resp *signInResp) GetToken() string {
	return resp.token
}

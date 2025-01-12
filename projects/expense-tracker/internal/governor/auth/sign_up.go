package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/auth"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/validation"
)

func (r *Auth) SignUp(ctx context.Context, req controller.SignUpReq) (controller.SignUpResp, error) {
	log := r.logger.With(slog.String("handler", "SignUp"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	if err := validateSignUp(req); err != nil {
		log.WarnContext(ctx, "invalid fields", slog.Any("error", err))
		return nil, err
	}

	passHash, salt, err := auth.HashPassword(req.GetPassword(), r.conf.API.Pepper)
	if err != nil {
		log.ErrorContext(ctx, "hashing password failed", slog.Any("error", err))
		return nil, fmt.Errorf("hashing password failed %w", err)
	}

	dbReq := newCreateUserDBReq(req.GetEmail(), passHash, salt)
	dbResp, err := r.db.CreateUser(ctx, dbReq)
	if err != nil || dbResp == nil {
		log.ErrorContext(ctx, "create user db failed", slog.Any("error", err))
		return nil, fmt.Errorf("create user db failed %w", err)
	}

	u := &auth.UserData{
		ID:    dbResp.GetID(),
		Email: dbResp.GetEmail(),
	}
	t, err := auth.GenerateToken(u, r.conf.API.TokenSecret)
	if err != nil {
		log.ErrorContext(ctx, "generate token failed", slog.Any("error", err))
		return nil, fmt.Errorf("generate token failed %w", err)
	}

	log.InfoContext(
		ctx,
		"success",
		slog.String("email", dbResp.GetEmail()),
	)
	return newSignUpResp(t), nil
}

func validateSignUp(req controller.SignUpReq) error {
	if len(req.GetEmail()) > 0 && len(req.GetPassword()) > 0 {
		return nil
	}
	return validation.NewError("invalid credentials")
}

type createUserDBReq struct {
	email        string
	passwordHash string
	salt         string
}

func newCreateUserDBReq(email, passwordHash, salt string) *createUserDBReq {
	return &createUserDBReq{
		email:        email,
		passwordHash: passwordHash,
		salt:         salt,
	}
}

func (req *createUserDBReq) GetEmail() string {
	return req.email
}

func (req *createUserDBReq) GetPasswordHash() string {
	return req.passwordHash
}

func (req *createUserDBReq) GetSalt() string {
	return req.salt
}

type signUpResp struct {
	token string
}

func newSignUpResp(token string) *signUpResp {
	return &signUpResp{token: token}
}

func (resp *signUpResp) GetToken() string {
	return resp.token
}

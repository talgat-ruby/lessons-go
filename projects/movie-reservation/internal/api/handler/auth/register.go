package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/auth"
	dbAuth "github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db/auth"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type RegisterRequest struct {
	Data *AuthUser `json:"data"`
}

type RegisterResponse struct {
	Data *AuthTokenPair `json:"data"`
}

func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "Register")

	// request parse
	requestBody := &RegisterRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// passwordHash and salt
	passHash, salt, err := auth.HashPassword(requestBody.Data.Password, os.Getenv("TOKEN_PEPPER"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"hashing password failed",
			"error", err,
		)
		http.Error(w, "hashing password failed", http.StatusInternalServerError)
		return
	}

	// db request
	userModel := &dbAuth.ModelUser{
		Email:        requestBody.Data.Email,
		PasswordHash: passHash,
		Salt:         salt,
	}
	dbResp, err := h.db.Register(ctx, &dbAuth.RegisterInput{userModel})
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to insert user data",
			"error", err,
		)
		http.Error(w, "failed to insert user data", http.StatusInternalServerError)
		return
	}

	// create token pair
	tokenPair, err := auth.GenerateTokenPair(
		&auth.UserData{
			ID:    fmt.Sprint(dbResp.ID),
			Email: dbResp.Email,
		},
		os.Getenv("TOKEN_SECRET"),
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("invalid request %w", err), http.StatusBadRequest)
		return
	}

	respBody := &RegisterResponse{
		Data: &AuthTokenPair{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}
	if err := response.JSON(
		w,
		http.StatusOK,
		respBody,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}

	log.InfoContext(
		ctx,
		"success register user",
		"user email", userModel.Email,
	)
	return
}

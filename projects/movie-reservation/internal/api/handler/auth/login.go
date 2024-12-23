package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/auth"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type LoginRequest struct {
	Data *AuthUser `json:"data"`
}

type LoginResponse struct {
	Data *AuthTokenPair `json:"data"`
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "Login")

	// request parse
	requestBody := &LoginRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// db request
	dbResp, err := h.db.Login(ctx, requestBody.Data.Email)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}

	if dbResp == nil {
		log.ErrorContext(
			ctx,
			"row is empty",
			"email", requestBody.Data.Email,
		)
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}

	// passwordHash and salt
	isValid, err := auth.VerifyPassword(
		requestBody.Data.Password,
		os.Getenv("TOKEN_PEPPER"),
		dbResp.PasswordHash,
		dbResp.Salt,
	)
	if err != nil {
		log.ErrorContext(
			ctx,
			"verify password failed",
			"error", err,
		)
		http.Error(w, "verify password failed", http.StatusInternalServerError)
		return
	}
	if !isValid {
		log.ErrorContext(
			ctx,
			"invalid credentials",
			"email", requestBody.Data.Email,
		)
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}

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

	// response
	respBody := &LoginResponse{
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
		"success login user",
		"email", dbResp.Email,
	)
	return
}

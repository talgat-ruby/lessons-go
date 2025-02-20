package users

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/rest/pkg/httperror"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/pkg/httputils/response"
)

func (h *Users) PostUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "PostUser")

	// request parse
	reqBody := &postUserReq{}
	if err := request.JSON(w, r, reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	dbResp, err := h.db.CreateUser(ctx, reqBody)
	if err != nil || dbResp == nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid credentials", "", "").
			HandleError(w, err)
		return
	}

	respBody := &signInResp{
		Data: &signInRespData{
			ID:        dbResp.GetID(),
			Email:     dbResp.GetEmail(),
			CreatedAt: dbResp.GetCreatedAt(),
			UpdatedAt: dbResp.GetUpdatedAt(),
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
			slog.Any("error", err),
		)
		return
	}

	log.InfoContext(
		ctx,
		"success sign in user",
		slog.String("email", reqBody.GetEmail()),
	)
}

type postUserReq struct {
	Data *postUserReqData `json:"data"`
}

func (b *postUserReq) GetEmail() string {
	return b.Data.Email
}

func (b *postUserReq) GetPasswordHash() string {
	return b.Data.PasswordHash
}

func (b *postUserReq) GetSalt() string {
	return b.Data.Salt
}

type postUserReqData struct {
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Salt         string `json:"salt"`
}

type signInResp struct {
	Data *signInRespData `json:"data"`
}

type signInRespData struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

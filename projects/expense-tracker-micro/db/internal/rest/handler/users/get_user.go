package users

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/rest/pkg/httperror"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/pkg/httputils/response"
)

func (h *Users) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetUser")

	email := r.PathValue("email")

	dbReq := &dbFindUserByEmailReq{
		email: email,
	}

	dbResp, err := h.db.FindUserByEmail(ctx, dbReq)
	if err != nil || dbResp == nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid credentials", "", "").
			HandleError(w, err)
		return
	}

	respBody := &getUserResp{
		Data: &getUserRespData{
			ID:           dbResp.GetID(),
			Email:        dbResp.GetEmail(),
			PasswordHash: dbResp.GetEmail(),
			Salt:         dbResp.GetSalt(),
			CreatedAt:    dbResp.GetCreatedAt(),
			UpdatedAt:    dbResp.GetUpdatedAt(),
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
		"success get user",
		slog.String("email", email),
	)
}

type dbFindUserByEmailReq struct {
	email string
}

func (b *dbFindUserByEmailReq) GetEmail() string {
	return b.email
}

type getUserResp struct {
	Data *getUserRespData `json:"data"`
}

type getUserRespData struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"PasswordHash"`
	Salt         string    `json:"salt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

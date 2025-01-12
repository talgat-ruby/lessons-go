package auth

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/validation"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/response"
)

func (h *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "SignIn")

	// request parse
	reqBody := &signInReq{}
	if err := request.JSON(w, r, reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	ctrlResp, err := h.ctrl.SignIn(ctx, reqBody)
	if err != nil || ctrlResp == nil {
		valError := new(validation.Error)
		switch {
		case errors.As(err, &valError):
			log.ErrorContext(
				ctx,
				"validation failed",
				slog.Any("error", err),
			)
			http.Error(w, "invalid credentials", http.StatusBadRequest)
			return
		default:
			log.ErrorContext(
				ctx,
				"failed to sign in user",
				slog.Any("error", err),
			)
			http.Error(w, "failed to sign in user", http.StatusInternalServerError)
			return
		}
	}

	respBody := &signInResp{
		Data: &signInRespData{
			Token: ctrlResp.GetToken(),
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

type signInReq struct {
	Data *signInReqData `json:"data"`
}

func (b *signInReq) GetEmail() string {
	return b.Data.Email
}

func (b *signInReq) GetPassword() string {
	return b.Data.Password
}

type signInReqData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResp struct {
	Data *signInRespData `json:"data"`
}

type signInRespData struct {
	Token string `json:"token"`
}

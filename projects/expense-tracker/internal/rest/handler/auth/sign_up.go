package auth

import (
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/pkg/httperror"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/response"
)

func (h *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "SignUp")

	// request parse
	reqBody := new(signUpReq)
	if err := request.JSON(w, r, reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	ctrlResp, err := h.ctrl.SignUp(ctx, reqBody)
	if err != nil || ctrlResp == nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid credentials", "", "").
			HandleError(w, err)
		return
	}

	respBody := &signUpResp{
		Data: &signUpRespData{
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
		"success sign up user",
		slog.String("email", reqBody.GetEmail()),
	)
}

type signUpReq struct {
	Data *signUpReqData `json:"data"`
}

func (b *signUpReq) GetEmail() string {
	return b.Data.Email
}

func (b *signUpReq) GetPassword() string {
	return b.Data.Password
}

type signUpReqData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpResp struct {
	Data *signUpRespData `json:"data"`
}

type signUpRespData struct {
	Token string `json:"token"`
}

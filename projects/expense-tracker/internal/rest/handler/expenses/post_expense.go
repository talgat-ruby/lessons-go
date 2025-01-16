package expenses

import (
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/pkg/httperror"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/response"
)

func (h *Expenses) PostExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "PostExpense")

	// request parse
	reqBody := new(postExpenseReq)
	if err := request.JSON(w, r, reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	_, err := h.ctrl.NewExpense(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid values", "", "").
			HandleError(w, err)
		return
	}

	if err := response.JSON(
		w,
		http.StatusCreated,
		nil,
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
		"success",
	)
}

type postExpenseReq struct {
	Data *postExpenseReqData `json:"data"`
}

func (b *postExpenseReq) GetAmount() int64 {
	return b.Data.Amount
}

func (b *postExpenseReq) GetCategory() string {
	return b.Data.Category
}

type postExpenseReqData struct {
	Amount   int64  `json:"amount"`
	Category string `json:"category"`
}

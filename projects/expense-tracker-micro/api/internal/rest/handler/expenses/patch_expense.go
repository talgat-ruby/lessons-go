package expenses

import (
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/rest/pkg/httperror"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/pkg/httputils/response"
)

func (h *Expenses) PatchExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "PatchExpenses")

	id := r.PathValue("id")

	// request parse
	reqBody := new(patchExpenseReq)
	if err := request.JSON(w, r, reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}
	reqBody.id = id

	ctrlResp, err := h.ctrl.AlterExpense(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid values", "", "").
			HandleError(w, err)
		return
	}
	if ctrlResp == nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	if err := response.JSON(
		w,
		http.StatusNoContent,
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

type patchExpenseReq struct {
	id   string
	Data *patchExpensesReqData `json:"data"`
}

func (b *patchExpenseReq) GetID() string {
	return b.id
}

func (b *patchExpenseReq) GetAmount() *int64 {
	return b.Data.Amount
}

func (b *patchExpenseReq) GetCategory() *string {
	return b.Data.Category
}

type patchExpensesReqData struct {
	Amount   *int64  `json:"amount"`
	Category *string `json:"category"`
}

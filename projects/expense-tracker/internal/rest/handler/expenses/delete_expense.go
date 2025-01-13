package expenses

import (
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/pkg/httperror"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/response"
)

func (h *Expenses) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "DeleteExpenses")

	id := r.PathValue("id")

	// request parse
	reqBody := &deleteExpenseReq{
		id: id,
	}

	ctrlResp, err := h.ctrl.RemoveExpense(ctx, reqBody)
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

type deleteExpenseReq struct {
	id string
}

func (b *deleteExpenseReq) GetID() string {
	return b.id
}

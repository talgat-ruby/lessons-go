package expenses

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/validation"
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

	_, err := h.ctrl.RemoveExpense(ctx, reqBody)
	if err != nil {
		valError := new(validation.Error)
		switch {
		case errors.As(err, &valError):
			log.ErrorContext(
				ctx,
				"validation failed",
				slog.Any("error", err),
			)
			http.Error(w, "invalid values", http.StatusBadRequest)
			return
		default:
			log.ErrorContext(
				ctx,
				"fail from ctrl",
				slog.Any("error", err),
			)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
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

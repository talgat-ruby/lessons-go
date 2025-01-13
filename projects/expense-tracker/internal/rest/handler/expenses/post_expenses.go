package expenses

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/validation"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/response"
)

func (h *Expenses) PostExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "PostExpenses")

	// request parse
	reqBody := new(postExpensesReq)
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

type postExpensesReq struct {
	Data *postExpensesReqData `json:"data"`
}

func (b *postExpensesReq) GetAmount() int64 {
	return b.Data.Amount
}

func (b *postExpensesReq) GetCategory() string {
	return b.Data.Category
}

type postExpensesReqData struct {
	Amount   int64  `json:"amount"`
	Category string `json:"category"`
}

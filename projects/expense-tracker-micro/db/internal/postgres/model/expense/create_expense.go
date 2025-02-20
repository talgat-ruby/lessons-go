package expense

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/postgres/generated/expenses/public/model"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/types/database"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/validation"
)

func (m *Expense) CreateExpense(ctx context.Context, req database.CreateExpenseReq) (database.CreateExpenseResp, error) {
	log := m.logger.With(slog.String("handler", "CreateExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	userId, err := strconv.Atoi(req.GetUserID())
	if err != nil {
		log.ErrorContext(ctx, "could not convert id", slog.Any("error", err))
		return nil, validation.NewError("invalid user id")
	}

	mdl := model.Expense{
		UserID:   int32(userId),
		Amount:   req.GetAmount(),
		Category: mapStringToModelCategory(req.GetCategory()),
	}
	stmt := table.Expense.
		INSERT(table.Expense.UserID, table.Expense.Amount, table.Expense.Category).
		MODEL(mdl).
		RETURNING(table.Expense.AllColumns)
	if _, err := stmt.ExecContext(ctx, m.db); err != nil {
		switch {
		case errors.Is(err, qrm.ErrNoRows):
			log.WarnContext(ctx, "no rows found")
			return nil, nil
		default:
			log.ErrorContext(ctx, "failed to insert expense", slog.Any("error", err))
			return nil, fmt.Errorf("failed to insert expense: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}

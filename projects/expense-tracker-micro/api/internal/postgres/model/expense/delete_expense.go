package expense

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/database"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/validation"
)

func (m *Expense) DeleteExpense(ctx context.Context, req database.DeleteExpenseReq) (database.DeleteExpenseResp, error) {
	log := m.logger.With(slog.String("handler", "DeleteExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	id, err := strconv.Atoi(req.GetID())
	if err != nil {
		log.ErrorContext(ctx, "could not convert id", slog.Any("error", err))
		return nil, validation.NewError("invalid expense id")
	}

	userId, err := strconv.Atoi(req.GetUserID())
	if err != nil {
		log.ErrorContext(ctx, "could not convert id", slog.Any("error", err))
		return nil, validation.NewError("invalid user id")
	}

	stmt := table.Expense.
		DELETE().
		WHERE(
			pg.AND(
				table.Expense.ID.EQ(pg.Int(int64(id))),
				table.Expense.UserID.EQ(pg.Int(int64(userId))),
			),
		)
	if _, err := stmt.ExecContext(ctx, m.db); err != nil {
		switch {
		case errors.Is(err, qrm.ErrNoRows):
			log.WarnContext(ctx, "no rows found")
			return nil, nil
		default:
			log.ErrorContext(ctx, "failed to delete expense", slog.Any("error", err))
			return nil, fmt.Errorf("failed to delete expense: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}

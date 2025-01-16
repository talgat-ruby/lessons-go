package expense

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

func (m *Expense) ListExpense(ctx context.Context, req database.ListExpenseReq) (database.ListExpenseResp, error) {
	log := m.logger.With(slog.String("handler", "ListExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	// userId, err := strconv.Atoi(req.GetUserID())
	// if err != nil {
	// 	log.ErrorContext(ctx, "could not convert id", slog.Any("error", err))
	// 	return nil, validation.NewError("invalid user id")
	// }
	//
	// mdl := model.Expense{
	// 	UserID:   int32(userId),
	// 	Amount:   req.GetAmount(),
	// 	Category: mapStringToModelCategory(req.GetCategory()),
	// }
	// stmt := table.Expense.
	// 	INSERT(table.Expense.UserID, table.Expense.Amount, table.Expense.Category).
	// 	MODEL(mdl).
	// 	RETURNING(table.Expense.AllColumns)
	// if _, err := stmt.ExecContext(ctx, m.db); err != nil {
	// 	switch {
	// 	case errors.Is(err, qrm.ErrNoRows):
	// 		log.WarnContext(ctx, "no rows found")
	// 		return nil, nil
	// 	default:
	// 		log.ErrorContext(ctx, "failed to insert expense", slog.Any("error", err))
	// 		return nil, fmt.Errorf("failed to insert expense: %w", err)
	// 	}
	// }
	//
	// log.InfoContext(
	// 	ctx,
	// 	"success",
	// )
	return nil, nil
}

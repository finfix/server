package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/transaction/repository/model"
)

// CreateTransaction создает новую транзакцию
func (r *TransactionRepository) CreateTransaction(ctx context.Context, req model.CreateTransactionReq) (id uint32, err error) {

	// Создаем транзакцию
	return r.db.ExecWithLastInsertID(ctx, sq.Insert(`coin.transactions`).
		SetMap(map[string]any{
			"type_signatura":       req.Type,
			"date_transaction":     req.DateTransaction,
			"account_from_id":      req.AccountFromID,
			"account_to_id":        req.AccountToID,
			"amount_from":          req.AmountFrom,
			"amount_to":            req.AmountTo,
			"note":                 req.Note,
			"is_executed":          req.IsExecuted,
			"datetime_create":      req.DatetimeCreate,
			"created_by_user_id":   req.CreatedByUserID,
			"accounting_in_charts": req.AccountingInCharts,
		}),
	)
}

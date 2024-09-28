package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/transaction/repository/model"
	"server/internal/services/transaction/repository/transactionDDL"
)

// CreateTransaction создает новую транзакцию
func (r *TransactionRepository) CreateTransaction(ctx context.Context, req model.CreateTransactionReq) (id uint32, err error) {

	// Создаем транзакцию
	return r.db.ExecWithLastInsertID(ctx, sq.Insert(`coin.transactions`).
		SetMap(map[string]any{
			transactionDDL.ColumnType:               req.Type,
			transactionDDL.ColumnDate:               req.DateTransaction,
			transactionDDL.ColumnAccountFromID:      req.AccountFromID,
			transactionDDL.ColumnAccountToID:        req.AccountToID,
			transactionDDL.ColumnAmountFrom:         req.AmountFrom,
			transactionDDL.ColumnAmountTo:           req.AmountTo,
			transactionDDL.ColumnNote:               req.Note,
			transactionDDL.ColumnIsExecuted:         req.IsExecuted,
			transactionDDL.ColumnDatetimeCreate:     req.DatetimeCreate,
			transactionDDL.ColumnCreatedByUserID:    req.CreatedByUserID,
			transactionDDL.ColumnAccountingInCharts: req.AccountingInCharts,
		}),
	)
}

package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/transaction/model"
)

// GetTransactions возвращает все транзакции по фильтрам
func (r *TransactionRepository) GetTransactions(ctx context.Context, req model.GetTransactionsReq) (transactions []model.Transaction, err error) {

	q := sq.
		Select("t.*").
		From("coin.transactions t").
		Join("coin.accounts a1 ON a1.id = t.account_from_id").
		Join("coin.accounts a2 ON a2.id = t.account_to_id")

	// Добавляем фильтры
	if len(req.AccountGroupIDs) != 0 {
		q = q.Where(sq.Eq{
			"a1.account_group_id": req.AccountGroupIDs,
			"a2.account_group_id": req.AccountGroupIDs,
		})
	}
	if len(req.IDs) != 0 {
		q = q.Where(sq.Eq{"t.id": req.IDs})
	}
	if req.AccountID != nil {
		q = q.Where(sq.Or{
			sq.Eq{"t.account_from_id": *req.AccountID},
			sq.Eq{"t.account_to_id": *req.AccountID},
		})
	}
	if req.Type != nil {
		q = q.Where(sq.Eq{"t.type_signatura": *req.Type})
	}
	if req.DateFrom != nil {
		q = q.Where(sq.GtOrEq{"t.date_transaction": req.DateFrom})
	}
	if req.DateTo != nil {
		q = q.Where(sq.LtOrEq{"t.date_transaction": req.DateTo})
	}

	q = q.OrderBy("t.date_transaction DESC, t.datetime_create DESC")

	if req.Limit != nil {
		q = q.Limit(uint64(*req.Limit))
	}
	if req.Offset != nil {
		q = q.Offset(uint64(*req.Offset))
	}

	// Получаем транзакции
	if err = r.db.Select(ctx, &transactions, q); err != nil {
		return nil, err
	}

	return transactions, nil
}

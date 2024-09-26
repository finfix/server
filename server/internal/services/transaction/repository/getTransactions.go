package repository

import (
	"context"
	"fmt"
	"strings"

	"pkg/errors"

	"server/internal/services/transaction/model"
)

// GetTransactions возвращает все транзакции по фильтрам
func (repo *TransactionRepository) GetTransactions(ctx context.Context, req model.GetTransactionsReq) (transactions []model.Transaction, err error) {

	var (
		args        []any
		queryFields []string
	)

	// Добавляем фильтры
	if len(req.AccountGroupIDs) != 0 {
		_query, _args, err := repo.db.In(`account_group_id IN (?)`, req.AccountGroupIDs)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, fmt.Sprintf(`(a1.%v OR a2.%v)`, _query, _query))
		args = append(args, _args...)
		args = append(args, _args...)
	}
	if len(req.IDs) != 0 {
		_query, _args, err := repo.db.In(`t.id IN (?)`, req.IDs)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}
	if req.AccountID != nil {
		queryFields = append(queryFields, `(a1.id = ? OR a2.id = ?)`)
		args = append(args, *req.AccountID, *req.AccountID)
	}
	if req.Type != nil {
		queryFields = append(queryFields, `t.type_signatura = ?`)
		args = append(args, *req.Type)
	}
	if req.DateFrom != nil {
		queryFields = append(queryFields, `t.date_transaction >= ?`)
		args = append(args, req.DateFrom)
	}
	if req.DateTo != nil {
		queryFields = append(queryFields, `t.date_transaction < ?`)
		args = append(args, req.DateTo)
	}

	// Конструируем запрос
	query := fmt.Sprintf(`SELECT t.*
		   FROM coin.transactions t
			 JOIN coin.accounts a1 ON a1.id = t.account_from_id
			 JOIN coin.accounts a2 ON a2.id = t.account_to_id
		   WHERE %v
           ORDER BY 
             t.date_transaction DESC,
             t.datetime_create DESC`,
		strings.Join(queryFields, " AND "),
	)

	if req.Limit != nil {
		query += ` LIMIT ?`
		args = append(args, *req.Limit)
	}
	if req.Offset != nil {
		query += ` OFFSET ?`
		args = append(args, *req.Offset)
	}

	// Получаем транзакции
	if err = repo.db.Select(ctx, &transactions, query, args...); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.ClientReject.New("HTTP connection terminated")
		}
		return nil, err
	}

	return transactions, nil
}

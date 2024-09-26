package repository

import (
	"context"
	"fmt"
	"strings"

	"pkg/errors"

	"server/internal/services/transaction/model"
)

// UpdateTransaction редактирует транзакцию
func (repo *TransactionRepository) UpdateTransaction(ctx context.Context, fields model.UpdateTransactionReq) error {

	// Изменяем показатели транзакции
	var (
		args        []any
		queryFields []string
		query       string
	)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.IsExecuted != nil {
		queryFields = append(queryFields, `s_executed = ?`)
		args = append(args, fields.IsExecuted)
	}
	if fields.AccountFromID != nil {
		queryFields = append(queryFields, `account_from_id = ?`)
		args = append(args, fields.AccountFromID)
	}
	if fields.AccountToID != nil {
		queryFields = append(queryFields, `account_to_id = ?`)
		args = append(args, fields.AccountToID)
	}
	if fields.AmountFrom != nil {
		queryFields = append(queryFields, `amount_from = ?`)
		args = append(args, fields.AmountFrom)
	}
	if fields.AmountTo != nil {
		queryFields = append(queryFields, `amount_to = ?`)
		args = append(args, fields.AmountTo)
	}
	if fields.DateTransaction != nil {
		queryFields = append(queryFields, `date_transaction = ?`)
		args = append(args, fields.DateTransaction)
	}
	if fields.AccountingInCharts != nil {
		queryFields = append(queryFields, `accounting_in_charts = ?`)
		args = append(args, fields.AccountingInCharts)
	}
	if fields.Note != nil {
		if *fields.Note == "" {
			queryFields = append(queryFields, `note = NULL`)
		} else {
			queryFields = append(queryFields, `note = ?`)
			args = append(args, fields.Note)
		}
	}
	if len(queryFields) == 0 {
		if fields.TagIDs != nil {
			return nil
		}
		return errors.BadRequest.New("No fields to update")
	}

	// Конструируем запрос
	query = fmt.Sprintf(`
 			   UPDATE coin.transactions 
               SET %v
			   WHERE id = ?`,
		strings.Join(queryFields, ", "),
	)
	args = append(args, fields.ID)

	// Создаем транзакцию
	return repo.db.Exec(ctx, query, args...)
}

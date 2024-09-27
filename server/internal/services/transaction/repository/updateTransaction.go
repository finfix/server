package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"

	"server/internal/services/transaction/model"
)

// UpdateTransaction редактирует транзакцию
func (r *TransactionRepository) UpdateTransaction(ctx context.Context, fields model.UpdateTransactionReq) error {

	updates := make(map[string]any)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.IsExecuted != nil {
		updates["is_executed"] = *fields.IsExecuted
	}
	if fields.AccountFromID != nil {
		updates["account_from_id"] = *fields.AccountFromID
	}
	if fields.AccountToID != nil {
		updates["account_to_id"] = *fields.AccountToID
	}
	if fields.AmountFrom != nil {
		updates["amount_from"] = *fields.AmountFrom
	}
	if fields.AmountTo != nil {
		updates["amount_to"] = *fields.AmountTo
	}
	if fields.DateTransaction != nil {
		updates["date_transaction"] = *fields.DateTransaction
	}
	if fields.AccountingInCharts != nil {
		updates["accounting_in_charts"] = *fields.AccountingInCharts
	}
	if fields.Note != nil {
		updates["note"] = *fields.Note
	}

	// Проверяем, что есть поля для обновления
	if len(updates) == 0 {
		if fields.TagIDs != nil {
			return nil
		}
		return errors.BadRequest.New("No fields to update")
	}

	// Создаем транзакцию
	return r.db.Exec(ctx, sq.
		Update("coin.transactions").
		SetMap(updates).
		Where(sq.Eq{"id": fields.ID}),
	)
}

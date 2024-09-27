package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"

	accountRepoModel "server/internal/services/account/repository/model"
)

// UpdateAccount обновляет счет
func (r *AccountRepository) UpdateAccount(ctx context.Context, updateReqs map[uint32]accountRepoModel.UpdateAccountReq) error {

	for id, fields := range updateReqs {

		// Поля для обновления
		updates := make(map[string]any)

		// Добавляем в запрос только те поля, которые необходимо обновить
		if fields.IconID != nil {
			updates["icon_id"] = *fields.IconID
		}
		if fields.AccountingInHeader != nil {
			updates["accounting_in_header"] = *fields.AccountingInHeader
		}
		if fields.AccountingInCharts != nil {
			updates["accounting_in_charts"] = *fields.AccountingInCharts
		}
		if fields.Name != nil {
			updates["name"] = *fields.Name
		}
		if fields.Visible != nil {
			updates["visible"] = *fields.Visible
		}
		if fields.Budget.DaysOffset != nil {
			updates["budget_days_offset"] = *fields.Budget.DaysOffset
		}
		if fields.Budget.Amount != nil {
			updates["budget_amount"] = *fields.Budget.Amount
		}
		if fields.Budget.FixedSum != nil {
			updates["budget_fixed_sum"] = *fields.Budget.FixedSum
		}
		if fields.Budget.GradualFilling != nil {
			updates["budget_gradual_filling"] = *fields.Budget.GradualFilling
		}
		if fields.Currency != nil {
			updates["currency_signatura"] = *fields.Currency
		}
		if fields.SerialNumber != nil {
			updates["serial_number"] = *fields.SerialNumber
		}
		if fields.ParentAccountID != nil {
			if *fields.ParentAccountID == 0 {
				updates["parent_account_id"] = nil
			} else {
				updates["parent_account_id"] = *fields.ParentAccountID
			}
		}

		// Проверяем, переданы ли поля для обновления
		if len(updates) == 0 {
			if fields.Remainder == nil {
				return errors.BadRequest.New("No fields to update")
			}
			return nil
		}

		// Обновляем счет
		err := r.db.Exec(ctx, sq.
			Update("coin.accounts").
			SetMap(updates).
			Where(sq.Eq{"id": id}),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

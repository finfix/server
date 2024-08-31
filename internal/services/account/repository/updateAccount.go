package repository

import (
	"context"
	"fmt"
	"strings"

	accountRepoModel "server/internal/services/account/repository/model"
	"server/pkg/errors"
)

// UpdateAccount обновляет счет
func (repo *Repository) UpdateAccount(ctx context.Context, updateReqs map[uint32]accountRepoModel.UpdateAccountReq) error {

	for id, fields := range updateReqs {

		var (
			queryFields []string
			args        []any
		)

		// Добавляем в запрос только те поля, которые необходимо обновить
		if fields.IconID != nil {
			queryFields = append(queryFields, "icon_id = ?")
			args = append(args, fields.IconID)
		}
		if fields.AccountingInHeader != nil {
			queryFields = append(queryFields, "accounting_in_header = ?")
			args = append(args, fields.AccountingInHeader)
		}
		if fields.AccountingInCharts != nil {
			queryFields = append(queryFields, "accounting_in_charts = ?")
			args = append(args, fields.AccountingInCharts)
		}
		if fields.Name != nil {
			queryFields = append(queryFields, "name = ?")
			args = append(args, fields.Name)
		}
		if fields.Visible != nil {
			queryFields = append(queryFields, "visible = ?")
			args = append(args, fields.Visible)
		}
		if fields.Budget.DaysOffset != nil {
			queryFields = append(queryFields, "budget_days_offset = ?")
			args = append(args, fields.Budget.DaysOffset)
		}
		if fields.Budget.Amount != nil {
			queryFields = append(queryFields, "budget_amount = ?")
			args = append(args, fields.Budget.Amount)
		}
		if fields.Budget.FixedSum != nil {
			queryFields = append(queryFields, "budget_fixed_sum = ?")
			args = append(args, fields.Budget.FixedSum)
		}
		if fields.Budget.GradualFilling != nil {
			queryFields = append(queryFields, "budget_gradual_filling = ?")
			args = append(args, fields.Budget.GradualFilling)
		}
		if fields.Currency != nil {
			queryFields = append(queryFields, "currency_signatura = ?")
			args = append(args, fields.Currency)
		}
		if fields.SerialNumber != nil {
			queryFields = append(queryFields, "serial_number = ?")
			args = append(args, fields.SerialNumber)
		}
		if fields.ParentAccountID != nil {
			if *fields.ParentAccountID == 0 {
				queryFields = append(queryFields, "parent_account_id = NULL")
			} else {
				queryFields = append(queryFields, "parent_account_id = ?")
				args = append(args, fields.ParentAccountID)
			}
		}

		if len(queryFields) == 0 {
			if fields.Remainder == nil {
				return errors.BadRequest.New("No fields to update")
			}
			return nil
		}

		// Конструируем запрос
		query := fmt.Sprintf(`
				UPDATE coin.accounts 
				  SET %v 
				WHERE id = ?`,
			strings.Join(queryFields, ", "),
		)
		args = append(args, id)

		// Обновляем счета
		err := repo.db.Exec(ctx, query, args...)
		if err != nil {
			return err
		}
	}

	return nil
}

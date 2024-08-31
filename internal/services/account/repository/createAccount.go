package repository

import (
	"context"

	accountRepoModel "server/internal/services/account/repository/model"
)

// CreateAccount создает новый счет
func (repo *Repository) CreateAccount(ctx context.Context, account accountRepoModel.CreateAccountReq) (id uint32, serialNumber uint32, err error) {

	// Получаем текущий максимальный серийный номер
	row, err := repo.db.QueryRow(ctx, `
			SELECT COALESCE(MAX(serial_number), 1) AS serial_number
			FROM coin.accounts 
			WHERE account_group_id = ?`,
		account.AccountGroupID,
	)
	if err != nil {
		return id, serialNumber, err
	}
	if err = row.Scan(&serialNumber); err != nil {
		return id, serialNumber, err
	}
	serialNumber++

	// Создаем счет
	id, err = repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.accounts (
			  budget_amount,
			  name,
			  icon_id,
			  type_signatura,
			  currency_signatura,
			  visible,
			  account_group_id,
			  accounting_in_header,
			  accounting_in_charts,
			  budget_gradual_filling,
			  is_parent,
			  budget_fixed_sum,
			  budget_days_offset,        
			  parent_account_id,
			  created_by_user_id,
			  datetime_create,
			  serial_number
		  	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		account.Budget.Amount,
		account.Name,
		account.IconID,
		account.Type,
		account.Currency,
		account.Visible,
		account.AccountGroupID,
		account.AccountingInHeader,
		account.AccountingInCharts,
		account.Budget.GradualFilling,
		account.IsParent,
		account.Budget.FixedSum,
		account.Budget.DaysOffset,
		account.ParentAccountID,
		account.UserID,
		account.DatetimeCreate,
		serialNumber,
	)
	if err != nil {
		return id, serialNumber, err
	}
	return id, serialNumber, nil
}

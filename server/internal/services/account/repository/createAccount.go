package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	accountRepoModel "server/internal/services/account/repository/model"
)

// CreateAccount создает новый счет
func (r *AccountRepository) CreateAccount(ctx context.Context, account accountRepoModel.CreateAccountReq) (id uint32, serialNumber uint32, err error) {

	// Получаем максимальный серийный номер в группе счетов
	row, err := r.db.QueryRow(ctx, sq.
		Select("COALESCE(MAX(serial_number), 1) AS serial_number").
		From("coin.accounts").
		Where(sq.Eq{"account_group_id": account.AccountGroupID}),
	)
	if err != nil {
		return id, serialNumber, err
	}

	// Сканируем результат
	if err = row.Scan(&serialNumber); err != nil {
		return id, serialNumber, err
	}

	// Увеличиваем серийный номер для нового элемента
	serialNumber++

	// Создаем счет
	id, err = r.db.ExecWithLastInsertID(ctx, sq.
		Insert("coin.accounts").
		SetMap(map[string]any{
			"budget_amount":          account.Budget.Amount,
			"name":                   account.Name,
			"icon_id":                account.IconID,
			"type_signatura":         account.Type,
			"currency_signatura":     account.Currency,
			"visible":                account.Visible,
			"account_group_id":       account.AccountGroupID,
			"accounting_in_header":   account.AccountingInHeader,
			"accounting_in_charts":   account.AccountingInCharts,
			"budget_gradual_filling": account.Budget.GradualFilling,
			"is_parent":              account.IsParent,
			"budget_fixed_sum":       account.Budget.FixedSum,
			"budget_days_offset":     account.Budget.DaysOffset,
			"parent_account_id":      account.ParentAccountID,
			"created_by_user_id":     account.UserID,
			"datetime_create":        account.DatetimeCreate,
			"serial_number":          serialNumber,
		}))
	if err != nil {
		return id, serialNumber, err
	}

	return id, serialNumber, nil
}

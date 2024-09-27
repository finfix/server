package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
	"server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
)

// GetAccounts возвращает все счета, удовлетворяющие фильтрам
func (r *AccountRepository) GetAccounts(ctx context.Context, req accountRepoModel.GetAccountsReq) (accounts []model.Account, err error) {

	filters := make(sq.Eq)

	if len(req.AccountGroupIDs) != 0 {
		filters["a.account_group_id"] = req.AccountGroupIDs
	}
	if len(req.IDs) != 0 {
		filters["a.id"] = req.IDs
	}
	if len(req.Types) != 0 {
		filters["a.type_signatura"] = req.Types
	}
	if len(req.Currencies) != 0 {
		filters["a.currency_signatura"] = req.Currencies
	}
	if len(req.ParentAccountIDs) != 0 {
		filters["a.parent_account_id"] = req.ParentAccountIDs
	}
	if req.IsParent != nil {
		filters["a.is_parent"] = req.IsParent
	}
	if req.AccountingInHeader != nil {
		filters["a.accounting_in_header"] = req.AccountingInHeader
	}
	if req.AccountingInCharts != nil {
		filters["a.accounting_in_charts"] = req.AccountingInCharts
	}
	if req.Visible != nil {
		filters["a.visible"] = req.Visible
	}

	// Проверяем, что хоть один фильтр был передан
	if len(filters) == 0 {
		return accounts, errors.BadRequest.New("No filters")
	}

	// Выполняем запрос
	if err = r.db.Select(ctx, &accounts, sq.
		Select("a.*").
		From("coin.accounts a").
		Where(filters),
	); err != nil {
		return accounts, err
	}

	return accounts, nil
}

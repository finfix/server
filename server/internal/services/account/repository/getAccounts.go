package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	"pkg/errors"
	"server/internal/services/account/model"
	"server/internal/services/account/repository/accountDDL"
	accountRepoModel "server/internal/services/account/repository/model"
)

// GetAccounts возвращает все счета, удовлетворяющие фильтрам
func (r *AccountRepository) GetAccounts(ctx context.Context, req accountRepoModel.GetAccountsReq) (accounts []model.Account, err error) {

	filters := make(sq.Eq)

	if len(req.AccountGroupIDs) != 0 {
		filters[accountDDL.ColumnAccountGroupID] = req.AccountGroupIDs
	}
	if len(req.IDs) != 0 {
		filters[accountDDL.ColumnID] = req.IDs
	}
	if len(req.Types) != 0 {
		filters[accountDDL.ColumnType] = req.Types
	}
	if len(req.Currencies) != 0 {
		filters[accountDDL.ColumnCurrency] = req.Currencies
	}
	if len(req.ParentAccountIDs) != 0 {
		filters[accountDDL.ColumnParentAccountID] = req.ParentAccountIDs
	}
	if req.IsParent != nil {
		filters[accountDDL.ColumnIsParent] = req.IsParent
	}
	if req.AccountingInHeader != nil {
		filters[accountDDL.ColumnAccountingInHeader] = req.AccountingInHeader
	}
	if req.AccountingInCharts != nil {
		filters[accountDDL.ColumnAccountingInCharts] = req.AccountingInCharts
	}
	if req.Visible != nil {
		filters[accountDDL.ColumnVisible] = req.Visible
	}

	// Проверяем, что хоть один фильтр был передан
	if len(filters) == 0 {
		return accounts, errors.BadRequest.New("No filters")
	}

	// Выполняем запрос
	if err = r.db.Select(ctx, &accounts, sq.
		Select(ddlHelper.SelectAll).
		From(accountDDL.Table).
		Where(filters),
	); err != nil {
		return accounts, err
	}

	return accounts, nil
}

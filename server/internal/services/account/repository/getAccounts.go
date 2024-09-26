package repository

import (
	"context"
	"fmt"
	"strings"

	"pkg/errors"

	"server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
)

// GetAccounts возвращает все счета, удовлетворяющие фильтрам
func (repo *AccountRepository) GetAccounts(ctx context.Context, req accountRepoModel.GetAccountsReq) (accounts []model.Account, err error) {

	// Создаем конструктор запроса
	var (
		variables []any
		reqFields []string
	)

	if len(req.AccountGroupIDs) != 0 {
		_query, _args, err := repo.db.In(`a.account_group_id IN (?)`, req.AccountGroupIDs)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.IDs) != 0 {
		_query, _args, err := repo.db.In(`a.id IN (?)`, req.IDs)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.Types) != 0 {
		_query, _args, err := repo.db.In(`a.type_signatura IN (?)`, req.Types)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.Currencies) != 0 {
		_query, _args, err := repo.db.In(`a.currency_signatura IN (?)`, req.Currencies)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.ParentAccountIDs) != 0 {
		_query, _args, err := repo.db.In(`a.parent_account_id IN (?)`, req.ParentAccountIDs)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if req.IsParent != nil {
		reqFields = append(reqFields, `a.is_parent = ?`)
		variables = append(variables, req.IsParent)
	}
	if req.AccountingInHeader != nil {
		reqFields = append(reqFields, `a.accounting_in_header = ?`)
		variables = append(variables, req.AccountingInHeader)
	}
	if req.AccountingInCharts != nil {
		reqFields = append(reqFields, `a.accounting_in_charts = ?`)
		variables = append(variables, req.AccountingInCharts)
	}
	if req.Visible != nil {
		reqFields = append(reqFields, `a.visible = ?`)
		variables = append(variables, req.Visible)
	}

	if len(reqFields) == 0 {
		return accounts, errors.BadRequest.New("No filters")
	}

	// Выполняем запрос
	query := fmt.Sprintf(`
			SELECT a.*
			FROM coin.accounts a
			WHERE %v`, strings.Join(reqFields, " AND "))
	if err = repo.db.Select(ctx, &accounts, query, variables...); err != nil {
		return accounts, err
	}

	return accounts, nil
}

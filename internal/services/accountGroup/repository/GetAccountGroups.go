package repository

import (
	"context"
	"fmt"
	"strings"

	"server/pkg/errors"
	"server/internal/services/accountGroup/model"
)

func (repo *AccountGroupRepository) GetAccountGroups(ctx context.Context, filters model.GetAccountGroupsReq) (accountGroups []model.AccountGroup, err error) {

	var (
		queryArgs = []string{"ag.id != ?"}
		args      = []any{0}
	)

	if len(filters.AccountGroupIDs) != 0 {
		_queryArgs, _args, err := repo.db.In(`ag.id IN (?)`, filters.AccountGroupIDs)
		if err != nil {
			return accountGroups, err
		}
		queryArgs = append(queryArgs, _queryArgs)
		args = append(args, _args...)
	}

	if filters.Necessary.UserID != 0 {
		queryArgs = append(queryArgs, `utag.user_id = ?`)
		args = append(args, filters.Necessary.UserID)
	}

	if len(queryArgs) == 0 {
		return accountGroups, errors.BadRequest.New("No filters")
	}

	query := fmt.Sprintf(`
			SELECT ag.*
			FROM coin.account_groups ag
    		  JOIN coin.users_to_account_groups utag ON utag.account_group_id = ag.id
			WHERE %v`, strings.Join(queryArgs, " AND "))

	// Выполняем запрос
	if err = repo.db.Select(ctx, &accountGroups, query, args...); err != nil {
		return accountGroups, err
	}

	return accountGroups, nil
}

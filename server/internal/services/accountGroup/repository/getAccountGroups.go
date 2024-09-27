package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"

	"server/internal/services/accountGroup/model"
)

func (r *AccountGroupRepository) GetAccountGroups(ctx context.Context, req model.GetAccountGroupsReq) (accountGroups []model.AccountGroup, err error) {

	filtersEq := make(sq.Eq)

	if len(req.AccountGroupIDs) != 0 {
		filtersEq["ag.id"] = req.AccountGroupIDs
	}
	if req.Necessary.UserID != 0 {
		filtersEq["utag.user_id"] = req.Necessary.UserID
	}
	if len(filtersEq) == 0 {
		return accountGroups, errors.BadRequest.New("No req")
	}

	// Выполняем запрос
	return accountGroups, r.db.Select(ctx, &accountGroups, sq.
		Select("ag.*").
		From("coin.account_groups ag").
		Join("coin.users_to_account_groups utag ON utag.account_group_id = ag.id").
		Where(filtersEq),
	)
}

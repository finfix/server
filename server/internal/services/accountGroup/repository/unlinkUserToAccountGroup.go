package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (repo *AccountGroupRepository) UnlinkUserFromAccountGroup(ctx context.Context, userID, accountGroupID uint32) error {

	// Исполняем запрос на разрыв связи пользователя с группой счетов
	return repo.db.Exec(ctx, sq.
		Delete("coin.users_to_account_groups").
		Where(sq.And{
			sq.Eq{"user_id": userID},
			sq.Eq{"account_group_id": accountGroupID},
		}),
	)
}

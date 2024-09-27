package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (repo *AccountGroupRepository) LinkUserToAccountGroup(ctx context.Context, userID, accountGroupID uint32) error {

	// Исполняем запрос на связывание пользователя с группой счетов
	return repo.db.Exec(ctx, sq.
		Insert("coin.users_to_account_groups").
		SetMap(map[string]any{
			"user_id":          userID,
			"account_group_id": accountGroupID,
		}),
	)
}

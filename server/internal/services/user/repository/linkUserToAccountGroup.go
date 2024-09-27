package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *UserRepository) LinkUserToAccountGroup(ctx context.Context, userID uint32, accountGroupID uint32) error {
	return r.db.Exec(ctx, sq.
		Insert(`coin.users_to_account_groups`).
		SetMap(map[string]any{
			"user_id":          userID,
			"account_group_id": accountGroupID,
		}),
	)
}

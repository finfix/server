package repository

import "context"

func (repo *AccountGroupRepository) LinkUserToAccountGroup(ctx context.Context, userID, accountGroupID uint32) error {
	return repo.db.Exec(ctx, `
			INSERT INTO coin.users_to_account_groups (
			  user_id,
			  account_group_id
			) VALUES (?, ?)`,
		userID,
		accountGroupID,
	)
}

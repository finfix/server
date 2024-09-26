package repository

import "context"

func (repo *UserRepository) LinkUserToAccountGroup(ctx context.Context, userID uint32, accountGroupID uint32) error {
	return repo.db.Exec(ctx, `
			INSERT INTO coin.users_to_account_groups (
	          user_id,
	          account_group_id
	        ) VALUES (?, ?)`,
		userID,
		accountGroupID,
	)
}

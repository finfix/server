package repository

import "context"

func (repo *Repository) UnlinkUserFromAccountGroup(ctx context.Context, userID, accountGroupID uint32) error {
	return repo.db.Exec(ctx, `
			DELETE FROM coin.users_to_account_groups 
			WHERE user_id = ?
			  AND account_group_id = ?`,
		userID,
		accountGroupID,
	)
}

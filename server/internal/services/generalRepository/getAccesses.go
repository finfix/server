package generalRepository

import "context"

func (repo *GeneralRepository) getAccesses(ctx context.Context) (_ map[uint32]map[uint32]struct{}, err error) {

	usersToAccountsGroups := make(map[uint32]map[uint32]struct{})

	var result []struct {
		UserID         uint32 `db:"user_id"`
		AccountGroupID uint32 `db:"account_group_id"`
	}

	if err = repo.db.Select(ctx, &result, `
			SELECT u.id AS user_id, ag.id AS account_group_id
			FROM coin.account_groups ag
			JOIN coin.users_to_account_groups utag ON utag.account_group_id = ag.id 
			JOIN coin.users u ON utag.user_id = u.id`); err != nil {
		return nil, err
	}

	for _, item := range result {
		if _, ok := usersToAccountsGroups[item.UserID]; !ok {
			usersToAccountsGroups[item.UserID] = make(map[uint32]struct{})
		}
		usersToAccountsGroups[item.UserID][item.AccountGroupID] = struct{}{}
	}

	return usersToAccountsGroups, nil
}

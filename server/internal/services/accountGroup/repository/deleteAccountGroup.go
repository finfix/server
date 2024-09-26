package repository

import "context"

// DeleteAccountGroup удаляет группу счетов
func (repo *AccountGroupRepository) DeleteAccountGroup(ctx context.Context, id uint32) error {
	return repo.db.Exec(ctx, `DELETE FROM coin.account_groups WHERE id = ?`, id)
}

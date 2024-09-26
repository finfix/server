package repository

import "context"

// DeleteAccount удаляет счет
func (repo *AccountRepository) DeleteAccount(ctx context.Context, id uint32) error {
	return repo.db.Exec(ctx, `DELETE FROM coin.accounts WHERE id = ?`, id)
}

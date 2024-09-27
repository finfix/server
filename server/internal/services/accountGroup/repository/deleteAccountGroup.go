package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// DeleteAccountGroup удаляет группу счетов
func (repo *AccountGroupRepository) DeleteAccountGroup(ctx context.Context, id uint32) error {

	// Исполняем запрос на удаление группы счетов
	return repo.db.Exec(ctx, sq.
		Delete("coin.account_groups").
		Where(sq.Eq{"id": id}),
	)
}

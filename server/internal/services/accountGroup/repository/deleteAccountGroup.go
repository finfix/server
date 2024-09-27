package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// DeleteAccountGroup удаляет группу счетов
func (r *AccountGroupRepository) DeleteAccountGroup(ctx context.Context, id uint32) error {

	// Исполняем запрос на удаление группы счетов
	return r.db.Exec(ctx, sq.
		Delete("coin.account_groups").
		Where(sq.Eq{"id": id}),
	)
}

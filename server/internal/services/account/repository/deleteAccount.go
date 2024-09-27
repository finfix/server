package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// DeleteAccount удаляет счет
func (r *AccountRepository) DeleteAccount(ctx context.Context, id uint32) error {

	// Исполняем запрос на удаление счета
	return r.db.Exec(ctx, sq.
		Delete("coin.accounts").
		Where(sq.Eq{"id": id}),
	)
}

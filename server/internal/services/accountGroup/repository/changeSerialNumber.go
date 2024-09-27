package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// ChangeSerialNumbers вставляет группу счетов на новое место
func (r *AccountGroupRepository) ChangeSerialNumbers(ctx context.Context, oldValue, newValue uint32) error {

	// Формируем первичный запрос
	q := sq.Update("coin.account_groups")

	// Дополняем запрос в зависимости от того, в какую сторону двигаем элемент
	if newValue < oldValue {
		q = q.
			Set("serial_number", sq.Expr("serial_number + 1")).
			Where(sq.And{
				sq.GtOrEq{"serial_number": newValue},
				sq.Lt{"serial_number": oldValue},
			})
	} else {
		q = q.
			Set("serial_number", sq.Expr("serial_number - 1")).
			Where(sq.And{
				sq.Gt{"serial_number": oldValue},
				sq.LtOrEq{"serial_number": newValue},
			})
	}

	// Выполняем запрос
	return r.db.Exec(ctx, q)
}

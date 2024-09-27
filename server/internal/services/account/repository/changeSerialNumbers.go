package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *AccountRepository) ChangeSerialNumbers(ctx context.Context, accountGroupID, oldValue, newValue uint32) error {

	// Формируем первичный запрос
	q := sq.
		Update("coin.accounts")

	// Дополняем запрос в зависимости от того, в какую сторону двигаем элемент
	if newValue < oldValue {
		q = q.
			Set("serial_number", sq.Expr("serial_number + 1")).
			Where(sq.And{
				sq.Eq{"account_group_id": accountGroupID},
				sq.GtOrEq{"serial_number": newValue},
				sq.Lt{"serial_number": oldValue},
			})
	} else {
		q = q.
			Set("serial_number", sq.Expr("serial_number - 1")).
			Where(sq.And{
				sq.Eq{"account_group_id": accountGroupID},
				sq.Gt{"serial_number": oldValue},
				sq.LtOrEq{"serial_number": newValue},
			})
	}

	// Выполняем запрос
	return r.db.Exec(ctx, q)
}

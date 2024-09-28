package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	"server/internal/services/account/repository/accountDDL"
)

func (r *AccountRepository) ChangeSerialNumbers(ctx context.Context, accountGroupID, oldValue, newValue uint32) error {

	// Формируем первичный запрос
	q := sq.
		Update(accountDDL.Table)

	// Дополняем запрос в зависимости от того, в какую сторону двигаем элемент
	if newValue < oldValue {
		q = q.
			Set(accountDDL.ColumnSerialNumber, sq.Expr(ddlHelper.Plus(accountDDL.ColumnSerialNumber, 1))).
			Where(sq.And{
				sq.Eq{accountDDL.ColumnAccountGroupID: accountGroupID},
				sq.GtOrEq{accountDDL.ColumnSerialNumber: newValue},
				sq.Lt{accountDDL.ColumnSerialNumber: oldValue},
			})
	} else {
		q = q.
			Set(accountDDL.ColumnSerialNumber, sq.Expr(ddlHelper.Minus(accountDDL.ColumnSerialNumber, 1))).
			Where(sq.And{
				sq.Eq{accountDDL.ColumnAccountGroupID: accountGroupID},
				sq.Gt{accountDDL.ColumnSerialNumber: oldValue},
				sq.LtOrEq{accountDDL.ColumnSerialNumber: newValue},
			})
	}

	// Выполняем запрос
	return r.db.Exec(ctx, q)
}

package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
)

// CheckAccess проверяет, имеет ли набор групп счетов пользователя доступ к указанным идентификаторам счетов
func (r *AccountRepository) CheckAccess(ctx context.Context, accountGroupIDs, accountIDs []uint32) error {

	// Получаем все доступные счета по группам счетов и перечисленным счетам
	rows, err := r.db.Query(ctx, sq.
		Select("a.id").
		From("coin.accounts a").
		Where(sq.Eq{"a.account_group_id": accountGroupIDs}).
		Where(sq.Eq{"a.id": accountIDs}))
	if err != nil {
		return err
	}

	// Формируем мапу доступных счетов
	accessedAccountIDs := make(map[uint32]struct{})

	// Проходимся по каждому доступному счету
	for rows.Next() {

		// Считываем ID счета
		var accountID uint32
		if err = rows.Scan(&accountID); err != nil {
			return err
		}

		// Добавляем ID счета в мапу
		accessedAccountIDs[accountID] = struct{}{}
	}

	if len(accessedAccountIDs) == 0 {
		return errors.Forbidden.New("You don't have access to any of the requested accounts",
			errors.ParamsOption("AccountGroupIDs", accountGroupIDs, "AccountIDs", accountIDs),
		)
	}

	// Проходимся по каждому запрашиваемому счету
	for _, accountID := range accountIDs {

		// Если счета нет в мапе доступных счетов, возвращаем ошибку
		if _, ok := accessedAccountIDs[accountID]; !ok {
			return errors.Forbidden.New(fmt.Sprintf("You don't have access to account with ID %v", accountID),
				errors.ParamsOption("AccountGroupIDs", accountGroupIDs, "AccountID", accountID),
				errors.SkipPreviousCallerOption(),
			)
		}
	}

	return nil
}

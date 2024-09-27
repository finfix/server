package generalRepository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"

	"server/internal/services/generalRepository/checker"
)

// CheckUserAccessToObjects проверяет, имеет ли пользователь доступ к указанным идентификаторам объектов
func (repo *GeneralRepository) CheckUserAccessToObjects(ctx context.Context, checkType checker.CheckType, userID uint32, ids []uint32) error {

	accessedAccountGroupIDs := repo.GetAvailableAccountGroups(userID)

	if len(accessedAccountGroupIDs) == 0 {
		return errors.NotFound.New("Нет доступных объектов",
			errors.ParamsOption("UserID", userID, "IDs", ids, "Type", checkType),
		)
	}

	typeToWord := map[checker.CheckType]string{
		checker.Accounts:      "счетам",
		checker.AccountGroups: "группам счетов",
		checker.Transactions:  "транзакциям",
		checker.Tags:          "подкатегориям",
	}

	var q sq.SelectBuilder

	// В зависимости от проверяемого типа, выбираем таблицу
	switch checkType {

	case checker.Accounts:
		q = sq.
			Select("COUNT(*)").
			From("coin.accounts a").
			Where(sq.Eq{"a.account_group_id": accessedAccountGroupIDs}).
			Where(sq.Eq{"a.id": ids})

	case checker.AccountGroups:
		for _, accountGroupID := range ids {
			if _, ok := repo.accesses.Get()[userID][accountGroupID]; !ok {
				return errors.Forbidden.New("Access denied",
					errors.ParamsOption("UserID", userID, "IDs", ids, "Type", checkType),
					errors.HumanTextOption("Вы не имеете доступа к группе счетов с ID %v", accountGroupID),
					errors.SkipPreviousCallerOption(),
				)
			}
		}
		return nil

	case checker.Transactions:
		if len(ids) != 1 {
			return errors.InternalServer.New("Невозможно проверить доступ к нескольким транзакциям")
		}

		q = sq.
			Select("COUNT(*)").
			From("coin.transactions t").
			Join("coin.accounts a1 ON a1.id = t.account_from_id").
			Join("coin.accounts a2 ON a2.id = t.account_to_id").
			Where(sq.Eq{
				"a1.account_group_id": accessedAccountGroupIDs,
				"a2.account_group_id": accessedAccountGroupIDs,
				"t.id":                ids,
			})

	case checker.Tags:
		q = sq.
			Select("COUNT(*)").
			From("coin.tags t").
			Where(sq.Eq{
				"t.account_group_id": accessedAccountGroupIDs,
				"t.id":               ids,
			})
	}

	// Смотрим количество записей, которые удовлетворяют условию
	row, err := repo.db.QueryRow(ctx, q)
	if err != nil {
		return err
	}

	// Сканируем результат
	var countAccess uint32
	if err = row.Scan(&countAccess); err != nil {
		return err
	}

	// Если количество записей не равно количеству проверяемых идентификаторов, то возвращаем ошибку
	if countAccess != uint32(len(ids)) {
		return errors.Forbidden.New("Access denied",
			errors.ParamsOption("UserID", userID, "IDs", ids, "Type", checkType),
			errors.HumanTextOption("Вы не имеете доступа к %s", typeToWord[checkType]),
		)
	}

	return nil
}

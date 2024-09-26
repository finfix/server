package generalRepository

import (
	"context"
	"fmt"

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

	var (
		countAccess              uint32
		questionsAccountGroupIDs string
		argsAGs                  []any
		questionsIDs             string
		argsIDs                  []any
		args                     []any
		query                    string
	)

	// Добавляем в запрос условие на проверяемые идентификаторы групп счетов
	questionsAccountGroupIDs, argsAGs, err := repo.db.In("?", accessedAccountGroupIDs)
	if err != nil {
		return err
	}

	// Добавляем в запрос условие на проверяемые идентификаторы
	questionsIDs, argsIDs, err = repo.db.In("?", ids)
	if err != nil {
		return err
	}

	// В зависимости от проверяемого типа, выбираем таблицу
	switch checkType {

	case checker.Accounts:
		query = fmt.Sprintf(`
				SELECT COUNT(*)
				FROM coin.accounts a
				WHERE a.account_group_id IN (%v)
				AND a.id IN (%v)`,
			questionsAccountGroupIDs,
			questionsIDs,
		)
		args = append(args, argsAGs...)
		args = append(args, argsIDs...)

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
		query = fmt.Sprintf(`
				SELECT COUNT(*)
				FROM coin.transactions t 
				JOIN coin.accounts a1 ON a1.id = t.account_from_id 
				JOIN coin.accounts a2 ON a2.id = t.account_to_id
				WHERE a1.account_group_id IN (%v)
				AND a2.account_group_id IN (%v)
				AND t.id IN (%v)`,
			questionsAccountGroupIDs,
			questionsAccountGroupIDs,
			questionsIDs,
		)
		args = append(args, argsAGs...)
		args = append(args, argsAGs...)
		args = append(args, argsIDs...)

	case checker.Tags:
		query = fmt.Sprintf(`
				SELECT COUNT(*)
				FROM coin.tags t
				WHERE t.account_group_id IN (%v)
				AND t.id IN (%v)`,
			questionsAccountGroupIDs,
			questionsIDs,
		)
		args = append(args, argsAGs...)
		args = append(args, argsIDs...)
	}

	// Смотрим количество записей, которые удовлетворяют условию
	row, err := repo.db.QueryRow(ctx, query, args...)
	if err != nil {
		return err
	}
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

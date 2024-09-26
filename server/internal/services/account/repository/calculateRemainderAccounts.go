package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"

	accountRepoModel "server/internal/services/account/repository/model"
)

// CalculateRemainderAccounts возвращает остатки счетов
func (repo *AccountRepository) CalculateRemainderAccounts(ctx context.Context, req accountRepoModel.CalculateRemaindersAccountsReq) (map[uint32]decimal.Decimal, error) {

	var queryFields []string
	var args []any

	// Добавляем в запрос счета
	if len(req.IDs) != 0 {
		_query, _args, err := repo.db.In(
			`a.id IN (?)`, req.IDs,
		)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	// Добавляем в запрос типы счетов
	if len(req.Types) != 0 {
		_query, _args, err := repo.db.In(
			`a.type_signatura IN (?)`, req.Types,
		)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	// Добавляем в запрос группы счетов
	if len(req.AccountGroupIDs) != 0 {
		_query, _args, err := repo.db.In(
			`a.account_group_id IN (?)`, req.AccountGroupIDs,
		)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	// Добавляем в запрос даты
	if req.DateFrom != nil {
		queryFields = append(queryFields, `t.date_transaction >= ?`)
		args = append(args, req.DateFrom)
	}
	if req.DateTo != nil {
		queryFields = append(queryFields, `t.date_transaction < ?`)
		args = append(args, req.DateTo)
	}

	// Получаем сумму всех транзакций из счетов
	query := fmt.Sprintf(`
			SELECT t.account_to_id AS id, COALESCE(SUM(t.amount_to), 0) AS remainder
			FROM coin.transactions t
			JOIN coin.accounts a ON t.account_to_id = a.id
			WHERE %v
			GROUP BY t.account_to_id`,
		strings.Join(queryFields, " AND "))

	var amountsArray []struct {
		ID        uint32          `db:"id"`
		Remainder decimal.Decimal `db:"remainder"`
	}

	// Вычисляем сумму всех транзакций из счетов, id - сумма из
	if err := repo.db.Select(ctx, &amountsArray, query, args...); err != nil {
		return nil, err
	}

	// Формируем мапу с суммой исходящих транзакций в виде map[accountID]amountTransactions
	amountFromAccount := make(map[uint32]decimal.Decimal, len(amountsArray))
	for _, remainder := range amountsArray {
		amountFromAccount[remainder.ID] = remainder.Remainder
	}

	// Получаем сумму всех транзакций в счета
	query = fmt.Sprintf(`
			SELECT t.account_from_id AS id, COALESCE(SUM(t.amount_from), 0) AS remainder
			FROM coin.transactions t
			JOIN coin.accounts a ON t.account_from_id = a.id
			WHERE %v
			GROUP BY t.account_from_id`,
		strings.Join(queryFields, " AND "))

	// Вычисляем сумму всех транзакций в счета, id - сумма в
	if err := repo.db.Select(ctx, &amountsArray, query, args...); err != nil {
		return nil, err
	}

	// Формируем мапу с суммой входящих транзакций в виде map[accountID]amountTransactions
	amountToAccount := make(map[uint32]decimal.Decimal, len(amountFromAccount))
	for _, remainder := range amountsArray {
		amountToAccount[remainder.ID] = remainder.Remainder
	}

	// Проходим по всем счетам и вычисляем остаток разницей суммы в и из счета, формируем новую мапу
	amountMapToAccountID := make(map[uint32]decimal.Decimal)
	for id := range amountToAccount {
		amountMapToAccountID[id] = amountFromAccount[id].Sub(amountToAccount[id])
	}

	// Если счета нет в списке транзакций, то добавляем его со значением -сумма из счета
	for id := range amountFromAccount {
		if _, ok := amountMapToAccountID[id]; !ok {
			amountMapToAccountID[id] = amountFromAccount[id]
		}
	}

	return amountMapToAccountID, nil
}

package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/shopspring/decimal"

	accountRepoModel "server/internal/services/account/repository/model"
)

func applyFilters(q sq.SelectBuilder, req accountRepoModel.CalculateRemaindersAccountsReq) sq.SelectBuilder {
	if len(req.IDs) != 0 {
		q = q.Where(sq.Eq{"a.id": req.IDs})
	}
	if len(req.Types) != 0 {
		q = q.Where(sq.Eq{"a.type_signatura": req.Types})
	}
	if len(req.AccountGroupIDs) != 0 {
		q = q.Where(sq.Eq{"a.account_group_id": req.AccountGroupIDs})
	}
	if req.DateFrom != nil {
		q = q.Where(sq.GtOrEq{"t.date_transaction": req.DateFrom})
	}
	if req.DateTo != nil {
		q = q.Where(sq.LtOrEq{"t.date_transaction": req.DateTo})
	}
	return q
}

var amountsArray []struct {
	ID        uint32          `db:"id"`
	Remainder decimal.Decimal `db:"remainder"`
}

// CalculateRemainderAccounts возвращает остатки счетов
func (r *AccountRepository) CalculateRemainderAccounts(ctx context.Context, req accountRepoModel.CalculateRemaindersAccountsReq) (map[uint32]decimal.Decimal, error) {

	// Вычисляем сумму всех транзакций из счетов, id - сумма из
	if err := r.db.Select(ctx, &amountsArray,
		applyFilters(
			sq.
				Select("t.account_to_id AS id", "COALESCE(SUM(t.amount_from), 0) AS remainder").
				From("coin.transactions t").
				Join("coin.accounts a ON t.account_to_id = a.id").
				GroupBy("t.account_to_id"),
			req,
		),
	); err != nil {
		return nil, err
	}

	// Формируем мапу с суммой исходящих транзакций в виде map[accountID]amountTransactions
	amountFromAccount := make(map[uint32]decimal.Decimal, len(amountsArray))
	for _, remainder := range amountsArray {
		amountFromAccount[remainder.ID] = remainder.Remainder
	}

	// Вычисляем сумму всех транзакций в счета, id - сумма в
	if err := r.db.Select(ctx, &amountsArray,
		applyFilters(
			sq.
				Select("t.account_from_id AS id", "COALESCE(SUM(t.amount_from), 0) AS remainder").
				From("coin.transactions t").
				Join("coin.accounts a ON t.account_from_id = a.id").
				GroupBy("t.account_from_id"),
			req,
		),
	); err != nil {
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

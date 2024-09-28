package service

import (
	"context"

	"github.com/shopspring/decimal"

	"pkg/errors"

	"server/internal/services/account/model"
	"server/internal/services/account/model/accountType"
	accountRepoModel "server/internal/services/account/repository/model"
)

func (s *AccountService) calculateRemainders(ctx context.Context, filters model.GetAccountsReq) (map[uint32]decimal.Decimal, error) {

	// Считаем балансы обычных и долговых счетов
	calculatedRemainders, err := s.accountRepository.GetSumAllTransactionsToAccount(ctx, accountRepoModel.CalculateRemaindersAccountsReq{ //nolint:exhaustruct
		AccountGroupIDs: filters.AccountGroupIDs,
		Types: []accountType.Type{
			accountType.Debt,
			accountType.Regular,
		},
	})
	if err != nil {
		return nil, err
	}

	// Если тип счета расход, доход или балансировочный (или типа нет), то обязательно должен быть указан интервал дат
	if filters.Type == nil || *filters.Type == accountType.Earnings || *filters.Type == accountType.Expense || *filters.Type == accountType.Balancing {

		if filters.DateFrom == nil || filters.DateTo == nil {
			return nil, errors.BadRequest.New("dateFrom and dateTo must be specified")
		}

		// Считаем расходы и доходы за указанный период или даты
		earnAndExp, err := s.accountRepository.GetSumAllTransactionsToAccount(ctx, accountRepoModel.CalculateRemaindersAccountsReq{ //nolint:exhaustruct
			AccountGroupIDs: filters.AccountGroupIDs,
			Types: []accountType.Type{
				accountType.Earnings,
				accountType.Expense,
				accountType.Balancing,
			},
			DateFrom: filters.DateFrom,
			DateTo:   filters.DateTo,
		})
		if err != nil {
			return nil, err
		}

		// Добавляем балансы расходов, доходов и балансировочных счетов к остаткам
		for id, amount := range earnAndExp {
			calculatedRemainders[id] = amount
		}
	}
	return calculatedRemainders, err
}

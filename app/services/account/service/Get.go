package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository/checker"
)

// Get возвращает все счета, удовлетворяющие фильтрам
func (s *Service) Get(ctx context.Context, filters model.GetReq) (accounts []model.Account, err error) {

	// Проверяем доступ пользователя к группам счетов
	if len(filters.AccountGroupIDs) != 0 {
		if err = s.general.CheckAccess(ctx, checker.AccountGroups, filters.Necessary.UserID, filters.AccountGroupIDs); err != nil {
			return nil, err
		}
	} else {
		filters.AccountGroupIDs = s.general.GetAvailableAccountGroups(filters.Necessary.UserID)
	}

	// Получаем все счета
	accounts, err = s.accountRepository.Get(ctx, filters.ConvertToRepoReq())
	if err != nil {
		return nil, err
	}

	// Получаем остатки счетов
	calculatedRemainders, err := s.calculateRemainders(ctx, filters)
	if err != nil {
		return nil, err
	}

	// Заполняем остатки счетов
	for i, account := range accounts {
		if account.Type == accountType.Earnings || account.Type == accountType.Balancing {
			accounts[i].Remainder = -calculatedRemainders[account.ID]
		} else {
			accounts[i].Remainder = calculatedRemainders[account.ID]
		}
	}

	return accounts, nil
}

func (s *Service) calculateRemainders(ctx context.Context, filters model.GetReq) (map[uint32]float64, error) {

	// Считаем балансы обычных и долговых счетов
	calculatedRemainders, err := s.accountRepository.CalculateRemainderAccounts(ctx, accountRepoModel.CalculateRemaindersAccountsReq{
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
		earnAndExp, err := s.accountRepository.CalculateRemainderAccounts(ctx, accountRepoModel.CalculateRemaindersAccountsReq{
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

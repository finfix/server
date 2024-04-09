package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/pkg/slice"
	model2 "server/app/services/account/model"
	"server/app/services/account/model/accountType"
	"server/app/services/generalRepository/checker"
)

// Get возвращает все счета, удовлетворяющие фильтрам
func (s *Service) Get(ctx context.Context, filters model2.GetReq) (accounts []model2.Account, err error) {

	// Проверяем доступ пользователя к группам счетов
	if len(filters.AccountGroupIDs) != 0 {
		if err = s.general.CheckAccess(ctx, checker.AccountGroups, filters.UserID, filters.AccountGroupIDs); err != nil {
			return nil, err
		}
	} else {
		filters.AccountGroupIDs = s.general.GetAvailableAccountGroups(filters.UserID)
	}

	// Получаем все счета
	accounts, err = s.accountRepository.Get(ctx, filters)
	if err != nil {
		return nil, err
	}

	// Получаем остатки счетов
	calculatedRemainders, err := s.calculateRemainders(ctx, filters)
	if err != nil {
		return nil, err
	}

	// Заполняем остатки счетов
	// TODO: Переписать с учетом балансировочных счетов для каждой группы счета
	for i, account := range accounts {
		accounts[i].Remainder = calculatedRemainders[account.ID]
	}

	if filters.DateFrom != nil && filters.DateTo != nil {
		balancingAccounts, err := s.calculateBalancing(ctx, filters)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, balancingAccounts...)
	}

	return accounts, nil
}

func (s *Service) calculateRemainders(ctx context.Context, filters model2.GetReq) (map[uint32]float64, error) {

	// Считаем балансы всех счетов
	calculatedRemainders, err := s.accountRepository.CalculateRemainderAccounts(ctx, filters.AccountGroupIDs, filters.DateTo)
	if err != nil {
		return nil, err
	}

	// Если тип счета расход или доход (или типа нет), то обязательно должен быть указан интервал дат
	if filters.Type == nil || *filters.Type == accountType.Earnings || *filters.Type == accountType.Expense {

		if filters.DateFrom == nil || filters.DateTo == nil {
			return nil, errors.BadRequest.New("dateFrom and dateTo must be specified")
		}

		// Считаем расходы и доходы за указанный период или даты
		earnAndExp, err := s.accountRepository.CalculateExpensesAndEarnings(ctx, filters.AccountGroupIDs, *filters.DateFrom, *filters.DateTo)
		if err != nil {
			return nil, err
		}

		// Затираем остатки счетов расходов и доходов новыми
		for id, amount := range earnAndExp {
			calculatedRemainders[id] = amount
		}
	}
	return calculatedRemainders, err
}

func (s *Service) calculateBalancing(ctx context.Context, filters model2.GetReq) ([]model2.Account, error) {

	// Получаем суммы транзакций, разбитые по группам счетов и валютам
	balancingAmount, err := s.accountRepository.CalculateBalancingAmount(ctx, filters.AccountGroupIDs, *filters.DateFrom, *filters.DateTo)
	if err != nil {
		return nil, err
	}

	// Получаем дефолтные валюты для групп счетов
	_accountGroups, err := s.accountRepository.GetAccountGroups(ctx, model2.GetAccountGroupsReq{AccountGroupIDs: filters.AccountGroupIDs})
	if err != nil {
		return nil, err
	}
	accountGroupsMap := slice.ToMap(_accountGroups, func(ag model2.AccountGroup) uint32 { return ag.ID })

	// Получаем курсы валют
	currencies, err := s.general.GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	// Мапа группа счетов - счет
	accountsMap := make(map[uint32]model2.Account)

	// Адская костылина
	// TODO: Исправить
	var currentID uint32 = 1000000000

	// Считаем балансы счетов
	for _, amount := range balancingAmount {

		// Получаем счет из мапы
		account := accountsMap[amount.AccountGroupID]

		relation := currencies[accountGroupsMap[amount.AccountGroupID].Currency] / currencies[amount.Currency]
		account.Remainder += amount.Amount * relation

		accountsMap[amount.AccountGroupID] = account
	}

	accounts := make([]model2.Account, 0, len(accountsMap))

loop:
	for accountGroupID, account := range accountsMap {

		var accType accountType.Type

		remainder := account.Remainder

		switch {
		case account.Remainder > 0:
			accType = accountType.Earnings
		case account.Remainder < 0:
			accType = accountType.Expense
			remainder *= -1
		default:
			continue loop
		}

		accounts = append(accounts, model2.Account{
			ID:              currentID,
			Remainder:       remainder,
			Name:            "Балансировочный",
			IconID:          0,
			Type:            accType,
			Currency:        accountGroupsMap[accountGroupID].Currency,
			Visible:         true,
			AccountGroupID:  accountGroupID,
			Accounting:      true,
			ParentAccountID: nil,
			SerialNumber:    currentID,
		})
		currentID++
	}

	return accounts, nil
}

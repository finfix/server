package service

import (
	"context"
	"time"

	"core/app/enum/accountType"
	"core/app/internal/services/account/model"
	"pkg/datetime/period"
	"pkg/errors"
	"pkg/slice"
)

func (s *Service) QuickStatistic(ctx context.Context, req model.QuickStatisticReq) ([]model.QuickStatistic, error) {

	// Получаем все счета пользователя
	accounting := true
	//period := period.Month
	dateFrom, dateTo := period.GetTimeInterval(time.Now(), period.Month)
	accounts, err := s.Get(ctx, model.GetReq{
		UserID:     req.UserID,
		Accounting: &accounting,
		DateFrom:   &dateFrom,
		DateTo:     &dateTo,
	})
	if err != nil {
		return nil, err
	}

	// Получаем курсы валют
	rates, err := s.general.GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	// Получаем пользователя
	accountGroups, err := s.account.GetAccountGroups(ctx, model.GetAccountGroupsReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	if len(accountGroups) == 0 {
		return nil, errors.NotFound.NewCtx("Доступных групп счетов не найдено", "UserID: %d", req.UserID)
	}
	accountGroupsMap := slice.SliceToMap(accountGroups, func(ag model.AccountGroup) uint32 { return ag.ID })

	// Делаем мапу accountGroupID - QuickStatisticRes
	statisticMap := make(map[uint32]model.QuickStatistic)

	// Проходимся по каждому аккаунту
	for _, account := range accounts {
		if (
			// Если бюджеты и остатки пустые
			account.Budget == 0 && account.Remainder == 0) ||
			// Или тип счета доход
			account.Type == accountType.Earnings {
			// То статистику не собираем
			continue
		}
		accountGroupStatistic := statisticMap[account.AccountGroupID]
		accountGroupForAccount := accountGroupsMap[account.AccountGroupID]

		accountGroupStatistic.Currency = accountGroupForAccount.Currency
		accountGroupStatistic.AccountGroupID = accountGroupForAccount.ID

		relation := rates[accountGroupForAccount.Currency] / rates[account.Currency]

		switch account.Type {
		// Если счет является расходом
		case accountType.Expense:
			accountGroupStatistic.TotalBudget += account.Budget * relation
			accountGroupStatistic.TotalExpense += account.Remainder * relation
		case accountType.Regular, accountType.Debt:
			accountGroupStatistic.TotalRemainder += account.Remainder * relation
		}
		statisticMap[account.AccountGroupID] = accountGroupStatistic
	}

	return slice.MapToSlice(statisticMap), nil
}

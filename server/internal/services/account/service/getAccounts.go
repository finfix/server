package service

import (
	"context"

	"pkg/errors"

	"server/internal/services/account/model"
	"server/internal/services/account/model/accountType"
)

// GetAccounts возвращает все счета, удовлетворяющие фильтрам
func (s *AccountService) GetAccounts(ctx context.Context, filters model.GetAccountsReq) (accounts []model.Account, err error) {

	// Если в фильтрах переданы группы счетов
	if len(filters.AccountGroupIDs) != 0 {

		// Проверяем доступ пользователя к группам счетов
		if err = s.accountGroupService.CheckAccess(ctx, filters.Necessary.UserID, filters.AccountGroupIDs); err != nil {
			return nil, err
		}
	} else {

		// Получаем доступные для пользователя группы счетов и добавляем их в фильтры
		filters.AccountGroupIDs, err = s.userService.GetAccessedAccountGroups(ctx, filters.Necessary.UserID)
		if err != nil {
			return nil, err
		}
		if len(filters.AccountGroupIDs) == 0 {
			return nil, errors.NotFound.New("У пользователя нет доступных групп счетов")
		}
	}

	// Получаем все счета
	accounts, err = s.accountRepository.GetAccounts(ctx, filters.ConvertToRepoReq())
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
			accounts[i].Remainder = calculatedRemainders[account.ID].Neg()
		} else {
			accounts[i].Remainder = calculatedRemainders[account.ID]
		}
	}

	return accounts, nil
}

package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository/checker"
)

func (s *Service) ValidateUpdateParentAccountID(ctx context.Context, account model.Account, parentAccountID, userID uint32) error {

	if err := s.general.CheckAccess(ctx, checker.Accounts, userID, []uint32{parentAccountID}); err != nil {
		return err
	}

	// Получаем родительский счет
	parentAccounts, err := s.accountRepository.Get(ctx, accountRepoModel.GetReq{IDs: []uint32{parentAccountID}})
	if err != nil {
		return err
	}
	if len(parentAccounts) == 0 {
		return errors.NotFound.New("Родительский счет не найден", errors.Options{
			Params: map[string]any{
				"accountID": parentAccountID,
			},
		})
	}
	parentAccount := parentAccounts[0]

	// Проверяем, что указанный счет является родительским
	if parentAccount.ID != parentAccountID {
		return errors.BadRequest.New("Указанный счет не является родительским", errors.Options{
			Params: map[string]any{
				"accountID": parentAccountID,
			},
		})
	}

	// Проверяем, что счета находятся в одной группе
	if account.AccountGroupID != parentAccount.AccountGroupID {
		return errors.BadRequest.New("Счета находятся в разных группах", errors.Options{
			Params: map[string]any{
				"childAccountID":       account.ID,
				"childAccountGroupID":  account.AccountGroupID,
				"parentAccountID":      parentAccount.ID,
				"parentAccountGroupID": parentAccount.AccountGroupID,
			},
		})
	}

	// Проверяем, что типы счетов совпадают
	if account.Type != parentAccount.Type {
		return errors.BadRequest.New("Типы счетов не совпадают", errors.Options{
			Params: map[string]any{
				"childAccountID":  account.ID,
				"childType":       account.Type,
				"parentAccountID": parentAccount.ID,
				"parentType":      parentAccount.Type,
			},
		})
	}

	return nil
}

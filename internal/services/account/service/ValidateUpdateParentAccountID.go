package service

import (
	"context"

	"server/pkg/errors"
	"server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/internal/services/generalRepository/checker"
)

func (s *AccountService) ValidateUpdateParentAccountID(ctx context.Context, account model.Account, parentAccountID, userID uint32) error {

	if account.IsParent {
		return errors.BadRequest.New("Счет уже является родительским",
			errors.ParamsOption("accountID", account.ID),
		)
	}

	if err := s.general.CheckUserAccessToObjects(ctx, checker.Accounts, userID, []uint32{parentAccountID}); err != nil {
		return err
	}

	// Получаем родительский счет
	parentAccounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{IDs: []uint32{parentAccountID}}) //nolint:exhaustruct
	if err != nil {
		return err
	}
	if len(parentAccounts) == 0 {
		return errors.NotFound.New("Родительский счет не найден",
			errors.ParamsOption("accountID", parentAccountID),
		)
	}
	parentAccount := parentAccounts[0]

	// Проверяем, что указанный счет является родительским
	if parentAccount.ID != parentAccountID {
		return errors.BadRequest.New("Указанный счет не является родительским",
			errors.ParamsOption("accountID", parentAccountID),
		)
	}

	// Проверяем, что счета находятся в одной группе
	if account.AccountGroupID != parentAccount.AccountGroupID {
		return errors.BadRequest.New("Счета находятся в разных группах",
			errors.ParamsOption(
				"childAccountID", account.ID,
				"childAccountGroupID", account.AccountGroupID,
				"parentAccountID", parentAccount.ID,
				"parentAccountGroupID", parentAccount.AccountGroupID,
			))
	}

	// Проверяем, что типы счетов совпадают
	if account.Type != parentAccount.Type {
		return errors.BadRequest.New("Типы счетов не совпадают",
			errors.ParamsOption(
				"childAccountID", account.ID,
				"childType", account.Type,
				"parentAccountID", parentAccount.ID,
				"parentType", parentAccount.Type,
			))
	}

	return nil
}

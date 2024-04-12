package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/pkg/pointer"
	"server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository/checker"
)

// Update обновляет счета по конкретным полям
func (s *Service) Update(ctx context.Context, updateReq model.UpdateReq) (res model.UpdateRes, err error) {

	repoUpdateReqs := make(map[uint32]accountRepoModel.UpdateReq)
	repoUpdateReqs[updateReq.ID] = updateReq.ConvertToRepoReq()

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, updateReq.Necessary.UserID, []uint32{updateReq.ID}); err != nil {
		return res, err
	}

	// Получаем счет
	accounts, err := s.accountRepository.Get(ctx, accountRepoModel.GetReq{IDs: []uint32{updateReq.ID}})
	if err != nil {
		return res, err
	}
	if len(accounts) == 0 {
		return res, errors.NotFound.New("Счет не найден", errors.Options{
			Params: map[string]any{
				"accountID": updateReq.ID,
			},
		})
	}
	account := accounts[0]

	// Проверяем, что входные данные не противоречат разрешениям
	if err = s.permissionsService.CheckPermissions(updateReq, s.permissionsService.GetPermissions(account)); err != nil {
		return res, err
	}

	// Проверяем, можно ли привязать счет к родительскому счету
	if updateReq.ParentAccountID != nil {

		// Если привязываем счет к родительскому счету
		if *updateReq.ParentAccountID != 0 {

			// Проверяем возможность привязки
			if err := s.accountService.ValidateUpdateParentAccountID(ctx, account, *updateReq.ParentAccountID, updateReq.Necessary.UserID); err != nil {
				return res, err
			}
			account.ParentAccountID = updateReq.ParentAccountID

		} else { // Если отвязываем счет от родительского счета
			account.ParentAccountID = nil
		}
	}

	// Получаем дочерние счета
	var childrenAccounts []model.Account
	if account.IsParent {
		childrenAccounts, err = s.accountRepository.Get(ctx, accountRepoModel.GetReq{ParentAccountIDs: []uint32{updateReq.ID}})
		if err != nil {
			return res, err
		}
	}

	// Получаем родительский счет
	var parentAccount *model.Account
	if account.ParentAccountID != nil {
		parentAccounts, err := s.accountRepository.Get(ctx, accountRepoModel.GetReq{IDs: []uint32{*account.ParentAccountID}})
		if err != nil {
			return res, err
		}
		parentAccount = &parentAccounts[0]
	}

	if updateReq.Accounting != nil {
		account.Accounting = *updateReq.Accounting
	}
	s.CheckAccountingLogic(
		repoUpdateReqs,
		account,
		childrenAccounts,
		parentAccount,
	)

	if updateReq.Visible != nil {
		account.Visible = *updateReq.Visible
	}
	s.CheckVisibleLogic(
		repoUpdateReqs,
		account,
		childrenAccounts,
		parentAccount,
	)

	return res, s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {
		err, res = s.update(ctxTx, account, repoUpdateReqs)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *Service) update(ctx context.Context, account model.Account, updateReqs map[uint32]accountRepoModel.UpdateReq) (err error, res model.UpdateRes) {

	// Если передан остаток, редактируем его
	if updateReqs[account.ID].Remainder != nil {
		if res, err = s.accountService.ChangeRemainder(ctx, account, *updateReqs[account.ID].Remainder); err != nil {
			return err, res
		}
	}

	// Редактируем счет
	return s.accountRepository.Update(ctx, updateReqs), res
}

func (s *Service) CheckAccountingLogic(
	repoUpdateReqs map[uint32]accountRepoModel.UpdateReq,
	mainAccount model.Account,
	childrenAccounts []model.Account,
	parentAccount *model.Account,
) map[uint32]accountRepoModel.UpdateReq {

	// Если значение родительского счета отрицательное, а у дочернего счета положительное
	if parentAccount != nil && !parentAccount.Accounting && mainAccount.Accounting {
		// То возникает логическая ошибка, поэтому родительский счет становится подсчитываемым
		requestParentAccount := repoUpdateReqs[parentAccount.ID]

		requestParentAccount.Accounting = pointer.Pointer(true)

		repoUpdateReqs[parentAccount.ID] = requestParentAccount
	}

	for _, childAccount := range childrenAccounts {

		if childAccount.Accounting && !mainAccount.Accounting {
			// То возникает логическая ошибка, поэтому значение дочернего счета становится отрицательным
			requestChildAccount := repoUpdateReqs[childAccount.ID]

			requestChildAccount.Accounting = pointer.Pointer(false)

			repoUpdateReqs[childAccount.ID] = requestChildAccount
		}

	}

	return repoUpdateReqs
}

func (s *Service) CheckVisibleLogic(
	repoUpdateReqs map[uint32]accountRepoModel.UpdateReq,
	mainAccount model.Account,
	childrenAccounts []model.Account,
	parentAccount *model.Account,
) map[uint32]accountRepoModel.UpdateReq {

	// Если значение родительского счета отрицательное, а у дочернего счета положительное
	if parentAccount != nil && !parentAccount.Visible && mainAccount.Visible {
		// То возникает логическая ошибка, поэтому родительский счет становится подсчитываемым
		requestParentAccount := repoUpdateReqs[parentAccount.ID]

		requestParentAccount.Visible = pointer.Pointer(true)

		repoUpdateReqs[parentAccount.ID] = requestParentAccount
	}

	// Если видимость счета отрицательная, а подсчитываемость положительного
	if !mainAccount.Visible && mainAccount.Accounting {
		// То возникает логическая ошибка, поэтому подсчитываемость деактивируется
		requestMainAccount := repoUpdateReqs[mainAccount.ID]

		requestMainAccount.Accounting = pointer.Pointer(false)

		repoUpdateReqs[mainAccount.ID] = requestMainAccount

	}

	// Если значения родительского счета меняется, то значения дочерних счетов меняются на такое же
	for _, childAccount := range childrenAccounts {
		requestChildAccount := repoUpdateReqs[childAccount.ID]

		requestChildAccount.Visible = pointer.Pointer(mainAccount.Visible)

		if !childAccount.Visible && childAccount.Accounting {
			requestChildAccount.Accounting = pointer.Pointer(false)
		}

		repoUpdateReqs[childAccount.ID] = requestChildAccount
	}

	return repoUpdateReqs
}

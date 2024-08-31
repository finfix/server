package utils

import (
	"server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/pkg/pointer"
)

func HandleAccountingInHeader(
	repoUpdateReqs map[uint32]accountRepoModel.UpdateAccountReq,
	mainAccount model.Account,
	childrenAccounts []model.Account,
	parentAccount *model.Account,
) map[uint32]accountRepoModel.UpdateAccountReq {

	// Если значение родительского счета отрицательное, а у дочернего счета положительное
	if parentAccount != nil && !parentAccount.AccountingInHeader && mainAccount.AccountingInHeader {
		// То возникает логическая ошибка, поэтому родительский счет становится подсчитываемым
		requestParentAccount := repoUpdateReqs[parentAccount.ID]

		requestParentAccount.AccountingInHeader = pointer.Pointer(true)

		repoUpdateReqs[parentAccount.ID] = requestParentAccount
	}

	for _, childAccount := range childrenAccounts {

		if childAccount.AccountingInHeader && !mainAccount.AccountingInHeader {
			// То возникает логическая ошибка, поэтому значение дочернего счета становится отрицательным
			requestChildAccount := repoUpdateReqs[childAccount.ID]

			requestChildAccount.AccountingInHeader = pointer.Pointer(false)

			repoUpdateReqs[childAccount.ID] = requestChildAccount
		}

	}

	return repoUpdateReqs
}

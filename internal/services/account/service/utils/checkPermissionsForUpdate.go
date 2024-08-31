package utils

import (
	accountModel "server/internal/services/account/model"
	accountPermissionsModel "server/internal/services/accountPermissions/model"
	"server/pkg/errors"
)

func CheckAccountPermissionsForUpdate(req accountModel.UpdateAccountReq, permissions accountPermissionsModel.AccountPermissions) error {

	if (req.Budget.DaysOffset != nil || req.Budget.Amount != nil || req.Budget.FixedSum != nil || req.Budget.GradualFilling != nil) && !permissions.UpdateBudget {
		return errors.Forbidden.New("Нельзя менять бюджет")
	}

	if req.Currency != nil && !permissions.UpdateCurrency {
		return errors.Forbidden.New("Нельзя менять валюту")
	}

	if req.Remainder != nil && !permissions.UpdateRemainder {
		return errors.Forbidden.New("Нельзя менять остаток")
	}

	return nil
}

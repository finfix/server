package accountPermissions

import (
	accountModel "server/internal/services/account/model"
	"server/pkg/errors"
)

func (s *AccountPermissionsService) CheckAccountPermissions(req accountModel.UpdateAccountReq, permissions AccountPermissions) error {

	if (req.Budget.DaysOffset != nil || req.Budget.Amount != nil || req.Budget.FixedSum != nil || req.Budget.GradualFilling != nil) && !permissions.UpdateBudget {
		return errors.BadRequest.New("Нельзя менять бюджет")
	}

	if req.Currency != nil && !permissions.UpdateCurrency {
		return errors.BadRequest.New("Нельзя менять валюту")
	}

	if req.Remainder != nil && !permissions.UpdateRemainder {
		return errors.BadRequest.New("Нельзя менять остаток")
	}

	return nil
}

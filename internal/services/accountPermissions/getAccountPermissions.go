package accountPermissions

import (
	"context"

	"server/internal/services/account/model"
	"server/internal/services/account/model/accountType"
)

func (s *AccountPermissionsService) GetAccountPermissions(account model.Account) AccountPermissions {
	typeToPermissions, isParentToPermissions := s.permissions.get()
	return joinAccountPermissions(
		generalPermissions,
		typeToPermissions[account.Type],
		isParentToPermissions[account.IsParent],
	)
}

func (s *AccountPermissionsService) getAccountPermissions(ctx context.Context) (
	_ map[accountType.Type]AccountPermissions,
	_ map[bool]AccountPermissions,
	err error,
) {

	rows, err := s.db.Query(ctx, `
		SELECT * 
		FROM permissions.account_permissions`)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	typeToPermissions := make(map[accountType.Type]AccountPermissions)
	isParentToPermissions := make(map[bool]AccountPermissions)

	for rows.Next() {
		var _accountType, actionType string
		var access bool
		if err := rows.Scan(&_accountType, &actionType, &access); err != nil {
			return nil, nil, err
		}

		var permission AccountPermissions
		switch _accountType {
		case "regular", "debt", "earnings", "expense", "balancing":
			permission = typeToPermissions[accountType.Type(_accountType)]
		case "parent", "general": //nolint:goconst
			permission = isParentToPermissions[_accountType == "parent"] //nolint:goconst
		}

		switch actionType {
		case "update_budget":
			permission.UpdateBudget = access
		case "update_remainder":
			permission.UpdateRemainder = access
		case "update_currency":
			permission.UpdateCurrency = access
		case "update_parent_account_id":
			permission.UpdateParentAccountID = access
		case "create_transaction":
			permission.CreateTransaction = access
		}

		switch _accountType {
		case "regular", "debt", "earnings", "expense", "balancing":
			typeToPermissions[accountType.Type(_accountType)] = permission
		case "parent", "general":
			isParentToPermissions[_accountType == "parent"] = permission
		}
	}

	return typeToPermissions, isParentToPermissions, nil
}

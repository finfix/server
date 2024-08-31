package accountPermissions

func joinAccountPermissions(permissions ...AccountPermissions) (joinedPermissions AccountPermissions) {
	joinedPermissions = generalPermissions
	for _, permission := range permissions {
		joinedPermissions.UpdateBudget = joinedPermissions.UpdateBudget && permission.UpdateBudget
		joinedPermissions.UpdateRemainder = joinedPermissions.UpdateRemainder && permission.UpdateRemainder
		joinedPermissions.UpdateCurrency = joinedPermissions.UpdateCurrency && permission.UpdateCurrency
		joinedPermissions.UpdateParentAccountID = joinedPermissions.UpdateParentAccountID && permission.UpdateParentAccountID

		joinedPermissions.CreateTransaction = joinedPermissions.CreateTransaction && permission.CreateTransaction
	}
	return joinedPermissions
}

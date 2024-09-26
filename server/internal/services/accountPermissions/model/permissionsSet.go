package model

import "server/internal/services/account/model/accountType"

type PermissionSet struct {
	TypeToPermissions     map[accountType.Type]AccountPermissions
	IsParentToPermissions map[bool]AccountPermissions
}

package repository

import (
	"context"

	"server/internal/services/account/model/accountType"
	"server/internal/services/accountPermissions/model"
	"server/pkg/log"
)

type permissionItem struct {
	AccountType string `db:"account_type"`
	ActionType  string `db:"action_type"`
	Access      bool   `db:"access"`
}

func (s *AccountPermissionsRepository) GetAccountPermissions(ctx context.Context) (permissionSet model.PermissionSet, err error) {

	// Проверяем наличие данных в кэше
	var ok bool
	if permissionSet, ok = s.cache.Get(struct{}{}); ok {
		return permissionSet, nil
	}

	log.Info(ctx, "Обновляем кэш с пермишенами на действия со счетами")

	var permissions []permissionItem

	// Делаем запрос в базу
	if err := s.db.Select(ctx, &permissions, `SELECT * FROM permissions.account_permissions`); err != nil {
		return permissionSet, err
	}

	// Инициализируем структуру
	permissionSet = model.PermissionSet{
		TypeToPermissions:     make(map[accountType.Type]model.AccountPermissions),
		IsParentToPermissions: make(map[bool]model.AccountPermissions),
	}

	// Проходимся по каждой строке из базы данных
	for _, permissionItem := range permissions {

		// Получаем права доступа, которые уже лежат в объекте permissionSet для данного типа аккаунта
		var permission model.AccountPermissions
		switch permissionItem.AccountType {
		case "regular", "debt", "earnings", "expense", "balancing":
			permission = permissionSet.TypeToPermissions[accountType.Type(permissionItem.AccountType)]
		case "parent", "general": //nolint:goconst
			permission = permissionSet.IsParentToPermissions[permissionItem.AccountType == "parent"] //nolint:goconst
		}

		// Смотрим на действие и присваиваем соответствующий доступ
		switch permissionItem.ActionType {
		case "update_budget":
			permission.UpdateBudget = permissionItem.Access
		case "update_remainder":
			permission.UpdateRemainder = permissionItem.Access
		case "update_currency":
			permission.UpdateCurrency = permissionItem.Access
		case "update_parent_account_id":
			permission.UpdateParentAccountID = permissionItem.Access
		case "create_transaction":
			permission.CreateTransaction = permissionItem.Access
		}

		// Сохраняем права доступа в объект permissionSet до следующей итерации
		switch permissionItem.AccountType {
		case "regular", "debt", "earnings", "expense", "balancing":
			permissionSet.TypeToPermissions[accountType.Type(permissionItem.AccountType)] = permission
		case "parent", "general":
			permissionSet.IsParentToPermissions[permissionItem.AccountType == "parent"] = permission
		}
	}

	// Сохраняем данные в кэш
	s.cache.Set(struct{}{}, permissionSet)

	// Возвращаем данные
	return permissionSet, nil
}

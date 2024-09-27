package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/log"
	"server/internal/services/user/model"
)

// GetAccessedAccountGroups возвращает доступы пользователей к группам счетов
func (r *UserRepository) GetAccessedAccountGroups(ctx context.Context, userID uint32) (accessedAccountGroupIDs []uint32, err error) {

	// Проверяем наличие данных в кэше
	var ok bool
	if accessedAccountGroupIDs, ok = r.accessedAccountGroupIDsCache.Get(userID); ok {
		return accessedAccountGroupIDs, nil
	}

	log.Info(ctx, "Обновляем кэш с доступами пользователей к группам счетов")

	var usersToAccountGroups []model.UserToAccountGroup

	// Получаем связки между пользователями и группами счетов
	if err = r.db.Select(ctx, &usersToAccountGroups, sq.
		Select("u.id AS user_id", "ag.id AS account_group_id").
		From("coin.account_groups ag").
		Join("coin.users_to_account_groups utag ON utag.account_group_id = ag.id").
		Join("coin.users u ON utag.user_id = u.id"),
	); err != nil {
		return nil, err
	}

	// Формируем мапу userID - []accountGroupID
	usersToAccountsGroups := make(map[uint32][]uint32)

	// Проходимся по каждой связке пользователя и группы счетов
	for _, userToAccountGroup := range usersToAccountGroups {
		accessedAccountGroups := usersToAccountsGroups[userToAccountGroup.UserID]
		accessedAccountGroups = append(accessedAccountGroups, userToAccountGroup.AccountGroupID)
		usersToAccountsGroups[userToAccountGroup.UserID] = accessedAccountGroups
	}

	// Проходимся по каждому пользователю и его группам счетов
	for userID, accountGroups := range usersToAccountsGroups {

		// Добавляем запись в кэш
		r.accessedAccountGroupIDsCache.Set(userID, accountGroups)
	}

	return usersToAccountsGroups[userID], nil
}

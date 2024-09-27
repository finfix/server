package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	userModel "server/internal/services/user/model"
)

// GetUsers Возвращает пользователей по фильтрам
func (r *UserRepository) GetUsers(ctx context.Context, filters userModel.GetUsersReq) (user []userModel.User, err error) {

	filtersEq := make(sq.Eq)

	if len(filters.IDs) > 0 {
		filtersEq["id"] = filters.IDs
	}
	if len(filters.Emails) > 0 {
		filtersEq["email"] = filters.Emails
	}

	return user, r.db.Select(ctx, &user, sq.
		Select("*").
		From("coin.users").
		Where(filtersEq),
	)
}

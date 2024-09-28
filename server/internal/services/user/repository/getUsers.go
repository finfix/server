package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	userModel "server/internal/services/user/model"
	"server/internal/services/user/repository/userDDL"
)

// GetUsers Возвращает пользователей по фильтрам
func (r *UserRepository) GetUsers(ctx context.Context, filters userModel.GetUsersReq) (user []userModel.User, err error) {

	filtersEq := make(sq.Eq)

	if len(filters.IDs) > 0 {
		filtersEq[userDDL.ColumnID] = filters.IDs
	}
	if len(filters.Emails) > 0 {
		filtersEq[userDDL.ColumnEmail] = filters.Emails
	}

	return user, r.db.Select(ctx, &user, sq.
		Select(ddlHelper.SelectAll).
		From(userDDL.Table).
		Where(filtersEq),
	)
}

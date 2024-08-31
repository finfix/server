package repository

import (
	"context"
	"fmt"
	"strings"

	userModel "server/internal/services/user/model"
)

// GetUsers Возвращает пользователей по фильтрам
func (repo *Repository) GetUsers(ctx context.Context, filters userModel.GetUsersReq) (user []userModel.User, err error) {

	query := `
			SELECT *
			FROM coin.users `

	var (
		queryArgs []string
		args      []any
	)

	if len(filters.IDs) > 0 {
		_query, _args, err := repo.db.In("id IN (?)", filters.IDs)
		if err != nil {
			return user, err
		}
		queryArgs = append(queryArgs, _query)
		args = append(args, _args...)
	}

	if len(filters.Emails) > 0 {
		_query, _args, err := repo.db.In("email IN (?)", filters.Emails)
		if err != nil {
			return user, err
		}
		queryArgs = append(queryArgs, _query)
		args = append(args, _args...)
	}

	if len(queryArgs) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(queryArgs, " AND "))
	}

	return user, repo.db.Select(ctx, &user, query, args...)
}

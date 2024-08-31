package repository

import (
	"context"
	"fmt"
	"strings"

	userRepoModel "server/internal/services/user/repository/model"
)

// UpdateUser редактирует пользователя
func (repo *Repository) UpdateUser(ctx context.Context, fields userRepoModel.UpdateUserReq) error {

	// Изменяем поля пользователя
	var (
		args        []any
		queryFields []string
		query       string
	)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.Email != nil {
		queryFields = append(queryFields, `email = ?`)
		args = append(args, fields.Email)
	}
	if fields.Name != nil {
		queryFields = append(queryFields, `name = ?`)
		args = append(args, fields.Name)
	}
	if fields.DefaultCurrency != nil {
		queryFields = append(queryFields, `default_currency_signatura = ?`)
		args = append(args, fields.DefaultCurrency)
	}
	if fields.PasswordHash != nil && fields.PasswordSalt != nil {
		queryFields = append(queryFields, `password_hash = ?`)
		args = append(args, fields.PasswordHash)
		queryFields = append(queryFields, `password_salt = ?`)
		args = append(args, fields.PasswordSalt)
	}

	if len(queryFields) == 0 {
		return nil
	}

	// Конструируем запрос
	query = fmt.Sprintf(`
 			   UPDATE coin.users 
               SET %v
			   WHERE id = ?`,
		strings.Join(queryFields, ", "),
	)
	args = append(args, fields.ID)

	// Обновляем пользователя
	return repo.db.Exec(ctx, query, args...)
}

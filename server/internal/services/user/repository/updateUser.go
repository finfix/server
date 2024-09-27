package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
	userRepoModel "server/internal/services/user/repository/model"
)

// UpdateUser редактирует пользователя
func (r *UserRepository) UpdateUser(ctx context.Context, fields userRepoModel.UpdateUserReq) error {

	// Изменяем поля пользователя
	updates := make(map[string]any)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.Email != nil {
		updates["email"] = *fields.Email
	}
	if fields.Name != nil {
		updates["name"] = *fields.Name
	}
	if fields.DefaultCurrency != nil {
		updates["default_currency_signatura"] = *fields.DefaultCurrency
	}
	if fields.PasswordHash != nil && fields.PasswordSalt != nil {
		updates["password_hash"] = *fields.PasswordHash
		updates["password_salt"] = *fields.PasswordSalt
	}

	if len(updates) == 0 {
		return errors.BadRequest.New("No fields to update")
	}

	// Обновляем пользователя
	return r.db.Exec(ctx, sq.
		Update("coin.users").
		SetMap(updates).
		Where(sq.Eq{"id": fields.ID}),
	)
}

package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	userModel "server/internal/services/user/model"
)

// CreateUser Создает нового пользователя
func (r *UserRepository) CreateUser(ctx context.Context, user userModel.CreateReq) (uint32, error) {

	return r.db.ExecWithLastInsertID(ctx, sq.
		Insert(`coin.users`).
		SetMap(map[string]any{
			"name":                       user.Name,
			"email":                      user.Email,
			"password_hash":              user.PasswordHash,
			"time_create":                user.TimeCreate,
			"default_currency_signatura": user.DefaultCurrency,
			"password_salt":              user.PasswordSalt,
			"is_admin":                   false,
		}),
	)
}

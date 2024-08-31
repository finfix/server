package repository

import (
	"context"

	userModel "server/internal/services/user/model"
)

// CreateUser Создает нового пользователя
func (repo *Repository) CreateUser(ctx context.Context, user userModel.CreateReq) (uint32, error) {

	return repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.users (
			  name, 
			  email, 
			  password_hash, 
			  time_create, 
			  default_currency_signatura,
			  password_salt,
			  is_admin
			) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.TimeCreate,
		user.DefaultCurrency,
		user.PasswordSalt,
		false,
	)
}

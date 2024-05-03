package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"server/app/pkg/sql"
	userModel "server/app/services/user/model"
)

type Repository struct {
	db sql.SQL
}

func (repo *Repository) LinkUserToAccountGroup(ctx context.Context, userID uint32, accountGroupID uint32) error {
	return repo.db.Exec(ctx, `
			INSERT INTO coin.users_to_account_groups (
	          user_id,
	          account_group_id
	        ) VALUES (?, ?)`,
		userID,
		accountGroupID)
}

// CreateUser Создает нового пользователя
func (repo *Repository) CreateUser(ctx context.Context, user userModel.CreateReq) (uint32, error) {

	return repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.users (
			  name, 
			  email, 
			  password_hash, 
			  time_create, 
			  verification_email_code,
			  fcm_token,
			  default_currency_signatura,
			  last_sync,
			  password_salt
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.TimeCreate,
		"",
		"",
		user.DefaultCurrency,
		time.Now(),
		user.PasswordSalt,
	)
}

// GetUsers Возвращает пользователей по фильтрам
func (repo *Repository) GetUsers(ctx context.Context, filters userModel.GetReq) (user []userModel.User, err error) {

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

func New(db sql.SQL, ) *Repository {
	return &Repository{
		db: db,
	}
}

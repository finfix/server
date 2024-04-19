package repository

import (
	"context"
	"fmt"
	"strings"

	"server/app/pkg/logging"
	"server/app/pkg/sql"
	userModel "server/app/services/user/model"
)

type Repository struct {
	db     sql.SQL
	logger *logging.Logger
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

// Create Создает нового пользователя
func (repo *Repository) CreateUser(ctx context.Context, user userModel.CreateReq) (uint32, error) {

	return repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.users (
			  name, 
			  email, 
			  password_hash, 
			  time_create, 
			  default_currency_signatura
			) VALUES (?, ?, ?, ?, ?)`,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.TimeCreate,
		user.DefaultCurrency)
}

// GetTransactions Возвращает пользователя по фильтрам
func (repo *Repository) GetTransactions(ctx context.Context, filters userModel.GetReq) (user []userModel.User, err error) {

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

func New(db sql.SQL, logger *logging.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

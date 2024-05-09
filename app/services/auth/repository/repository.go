package repository

import (
	"context"
	"time"

	"server/app/pkg/sql"
	authModel "server/app/services/auth/model"
)

// CreateSession Создает новую сессию для пользователя и добавляет ее в БД
func (repo *Repository) CreateSession(ctx context.Context, token string, timeExpiry time.Time, deviceID string, userID uint32) error {

	// Добавляем сессию в БД
	return repo.db.Exec(ctx, `
			INSERT INTO coin.sessions (
			  refresh_token, 
			  expires_at, 
			  device_id, 
			  user_id
        	) VALUES (?, ?, ?, ?)`,
		token,
		timeExpiry,
		deviceID,
		userID)
}

// DeleteSession Удаляет сессию пользователя
func (repo *Repository) DeleteSession(ctx context.Context, userID uint32, deviceID string) error {
	return repo.db.Exec(ctx, `
			DELETE 
			FROM coin.sessions 
			WHERE user_id = ? 
			  AND device_ID = ?`,
		userID,
		deviceID,
	)
}

// GetSession Возвращает сессию пользователя
func (repo *Repository) GetSession(ctx context.Context, req authModel.RefreshTokensReq) (session authModel.Session, err error) {
	return session, repo.db.Get(ctx, &session, `
			SELECT *
			FROM coin.sessions 
			WHERE user_id = ?
			  AND device_id = ?
			  AND refresh_token = ?
			LIMIT 1`,
		req.Necessary.UserID,
		req.Necessary.DeviceID,
		req.Token,
	)
}

type Repository struct {
	db sql.SQL
}

func New(
	db sql.SQL,

) *Repository {
	return &Repository{
		db: db,
	}
}

package repository

import (
	"context"
	"time"

	"server/app/config"
	"server/app/pkg/auth"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/sql"
)

// CreateSession Создает новую сессию для пользователя и добавляет ее в БД
func (repo *Repository) CreateSession(ctx context.Context, id uint32, deviceID string) (refreshToken string, err error) {

	// TODO: Будто бы это должно быть не тут
	// Создаем refresh token
	refreshToken, err = auth.NewRefreshToken()
	if err != nil {
		return "", err
	}

	// Получаем время жизни refresh token
	refreshDur, err := time.ParseDuration(config.GetConfig().Token.RefreshTokenTTL)
	if err != nil {
		return "", err
	}

	// Добавляем сессию в БД
	return refreshToken, repo.db.Exec(ctx, `
			INSERT INTO coin.sessions (
			  refresh_token, 
			  expires_at, 
			  device_id, 
			  user_id
        	) VALUES (?, ?, ?, ?)`,
		refreshToken,
		time.Now().Add(refreshDur),
		deviceID,
		id)
}

// DeleteSession Удаляет сессию пользователя
func (repo *Repository) DeleteSession(ctx context.Context, oldRefreshToken string) error {
	return repo.db.Exec(ctx, `
			DELETE coin.sessions 
			WHERE refresh_token = ?`,
		oldRefreshToken,
	)
}

// GetSession Возвращает сессию пользователя по refresh token
func (repo *Repository) GetSession(ctx context.Context, refreshToken string) (id uint32, deviceID string, err error) {

	var session struct {
		ExpiresAt time.Time `db:"expires_at"`
		ID        uint32    `db:"id"`
		DeviceID  string    `db:"device_id"`
	}

	// Получаем данные сессии
	err = repo.db.Get(ctx, &session, `
			SELECT user_id, expires_at, device_id 
			FROM coin.sessions 
			WHERE refresh_token = ? 
			LIMIT 1`,
		refreshToken,
	)
	if err != nil {
		if errors.As(err, sql.ErrNoRows) {
			return 0, "", errors.Unauthorized.New("Session not found", errors.Options{
				HumanText: "Необходимо пройти процедуру авторизации",
			})
		}
	}

	// Проверяем, не истекла ли сессия
	if session.ExpiresAt.Before(time.Now()) {
		return 0, "", errors.Unauthorized.New("Session ended", errors.Options{
			HumanText: "Истек строк действия токена авторизации, необходимо авторизоваться снова",
		})
	}

	return session.ID, session.DeviceID, nil
}

type Repository struct {
	db     sql.SQL
	logger *logging.Logger
}

func New(
	db sql.SQL,
	logger *logging.Logger,
) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}
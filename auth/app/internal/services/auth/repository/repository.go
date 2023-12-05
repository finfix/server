package repository

import (
	"context"
	"time"

	"logger/app/logging"
	"pkg/auth"
	"pkg/errors"
	"pkg/sql"

	"auth/app/internal/config"
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
		Id        uint32    `db:"id"`
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
			err = errors.Unauthorized.New("Session not found")
			return 0, "", errors.AddHumanText(err, "Необходимо пройти процедуру авторизации")
		}
	}

	// Проверяем, не истекла ли сессия
	if session.ExpiresAt.Before(time.Now()) {
		err = errors.Unauthorized.New("Session ended")
		return 0, "", errors.AddHumanText(err, "Истек строк действия токена авторизации, необходимо авторизоваться снова")
	}

	return session.Id, session.DeviceID, nil
}

type Repository struct {
	db     *sql.DB
	logger *logging.Logger
}

func New(
	db *sql.DB,
	logger *logging.Logger,
) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

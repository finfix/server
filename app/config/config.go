package config

import (
	"github.com/caarlos0/env/v7"

	"server/app/pkg/database/postgresql"
	"server/app/pkg/errors"
)

// Config - общая структура конфига
type Config struct {

	// Адрес для http-сервера
	HTTP string `env:"LISTEN_HTTP"`

	// Данные базы данных
	Repository postgresql.PostgreSQLConfig
	DBName     string `env:"DB_NAME"`

	// Информация для JWT-токенов
	Token struct {
		AccessTokenTTL  string `env:"AUTH_ACCESS_TOKEN_TTL"`
		RefreshTokenTTL string `env:"AUTH_REFRESH_TOKEN_TTL"`
		SigningKey      string `env:"AUTH_TOKEN_SIGNING_KEY"`
	}

	// Информация для шифрования паролей
	GeneralSalt string `env:"SHA_SALT"`

	// Ключи для работы с внешним API
	APIKeys struct {
		CurrencyProvider string `env:"API_KEY_CURRENCY_PROVIDER"`
	}

	// Доступы к телеграм-боту
	Telegram struct {
		Enabled bool   `env:"TG_BOT_ENABLED"`
		Token   string `env:"TG_BOT_TOKEN"`
		ChatID  int64  `env:"TG_CHAT_ID"`
	}

	Notifications struct {
		Enabled bool `env:"NOTIFICATIONS_ENABLED"`
		APNs    struct {
			TeamID      string `env:"NOTIFICATIONS_APNS_TEAM_ID"`
			KeyID       string `env:"NOTIFICATIONS_APNS_KEY_ID"`
			KeyFilePath string `env:"NOTIFICATIONS_APNS_KEY_FILE_PATH"`
		}
	}

	ServiceName string `env:"SERVICE_NAME"`
}

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, errors.InternalServer.Wrap(err)
	}
	return config, nil
}

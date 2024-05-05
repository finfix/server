package main

import (
	"github.com/caarlos0/env/v7"

	"server/app/pkg/database"
	"server/app/pkg/errors"
)

// Config - общая структура конфига
type config struct {

	// Адрес для http-сервера
	HTTP string `env:"LISTEN_HTTP"`

	// Ключ для админских методов
	AdminSecretKey string `env:"SECRET_KEY"`

	// Данные базы данных
	Repository database.RepoConfig
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
		Token  string `env:"TG_BOT_TOKEN"`
		ChatID int64  `env:"TG_CHAT_ID"`
	}
}

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() (config config, err error) {
	if err := env.Parse(&config); err != nil {
		return config, errors.InternalServer.Wrap(err)
	}
	return config, nil
}

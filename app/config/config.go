package config

import (
	"sync"

	"github.com/caarlos0/env/v7"

	"server/app/pkg/database"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
)

// Config - общая структура конфига
type Config struct {

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

var instance *Config
var once sync.Once

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			logging.GetLogger().Fatal(errors.InternalServer.Wrap(err))
		}
	})
	return instance
}

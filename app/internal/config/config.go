package config

import (
	"server/pkg/database"
	"sync"

	"server/pkg/errors"
	"server/pkg/logging"

	"github.com/caarlos0/env/v7"
)

// Общая структура конфига

type Config struct {
	HTTP      string `env:"LISTEN_HTTP"`
	SecretKey string `env:"SECRET_KEY"`

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
	SHASalt string `env:"SHA_SALT"`

	// Ключи для работы с внешним API
	ApiKeys struct {
		CurrencyProvider string `env:"API_KEY_CURRENCY_PROVIDER"`
	}

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

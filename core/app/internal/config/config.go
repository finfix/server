package config

import (
	"sync"

	"logger/app/logging"
	"pkg/database"
	"pkg/errors"

	"github.com/caarlos0/env/v7"
)

type Config struct {

	// Адреса связанных сервисов
	Services struct {
		Core struct {
			GRPC string `env:"CORE_LISTEN_GRPC"`
		}
		Logger struct {
			GRPC string `env:"LOGGER_LISTEN_GRPC"`
		}
	}

	// Данные базы данных
	Repository database.RepoConfig
	DBName     string `env:"CORE_DB_NAME"`

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
}

var instance *Config
var once sync.Once

// GetConfig returns a pointer to the Config struct
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			logging.GetLogger().Fatal(errors.InternalServer.Wrap(err))
		}
	})
	return instance
}

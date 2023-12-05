package config

import (
	"sync"

	"logger/app/logging"
	"pkg/errors"

	"github.com/caarlos0/env/v7"
)

// Общая структура конфига

type Config struct {
	Services struct {
		JSON struct {
			HTTP string `env:"JSON_LISTEN_HTTP"`
		}
		Core struct {
			GRPC string `env:"CORE_LISTEN_GRPC"`
		}
		Logger struct {
			GRPC string `env:"LOGGER_LISTEN_GRPC"`
		}
		Auth struct {
			GRPC string `env:"AUTH_LISTEN_GRPC"`
		}
	}
	SecretKey string `env:"SECRET_KEY"`
	Token     struct {
		SigningKey string `env:"AUTH_TOKEN_SIGNING_KEY"`
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

package config

import (
	"sync"

	"logger/app/logging"
	"pkg/database"
	"pkg/errors"

	"github.com/caarlos0/env/v7"
)

// Общая структура конфига

type Config struct {
	Services struct {
		Logger struct {
			GRPC string `env:"LOGGER_LISTEN_GRPC"`
		}
	}
	Repository database.RepoConfig
	DBName     string `env:"LOGGER_DB_NAME"`
	WorkDir    string `env:"WORKDIR"`
}

var instance *Config
var once sync.Once

// Функция чтения конфига
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			logging.GetLogger().Fatal(errors.InternalServer.Wrap(err))
		}
	})
	return instance
}

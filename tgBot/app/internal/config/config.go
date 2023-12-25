package config

import (
	"sync"

	"logger/app/logging"
	"pkg/errors"

	"github.com/caarlos0/env/v7"
)

type Config struct {

	// Адреса связанных сервисов
	Services struct {
		TgBot struct {
			GRPC string `env:"TGBOT_LISTEN_GRPC"`
		}
		Logger struct {
			GRPC string `env:"LOGGER_LISTEN_GRPC"`
		}
	}

	// Информация для JWT-токенов
	Telegram struct {
		Token  string `env:"TG_BOT_TOKEN"`
		ChatID int64  `env:"TG_CHAT_ID"`
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

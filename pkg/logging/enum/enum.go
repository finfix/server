package enum

import (
	"server/pkg/errors"
)

// Уровни логов
const (
	Warning = LogLevel("warning")
	Error   = LogLevel("error")
	Info    = LogLevel("info")
	Fatal   = LogLevel("fatal")
	Debug   = LogLevel("debug")
	Panic   = LogLevel("panic")
)

type LogLevel string

func LogLevelValidation(level string) error {
	switch LogLevel(level) {
	case Fatal, Info, Error, Warning, Debug, Panic:
	default:
		return errors.BadRequest.New("invalid level")
	}
	return nil
}

package log

import (
	"time"
)

// Log - Структура лога
type Log struct {
	Path    string
	Params  map[string]any
	Message string
	Level   logLevel
	Time    time.Time
	TaskID  *string
}

type uuidKeyType string

const uuidKey uuidKeyType = "uuid"

// loggerSettings - Конфигурация логгера
type loggerSettings struct {
	isOff bool
}

var logger = &loggerSettings{}

// Off выключает логгер
func Off() {
	logger.isOff = true
}

// init конфигурирует логгер
func init() {
	logger = &loggerSettings{}
}

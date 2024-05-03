package log

import (
	"time"
)

// Log - Структура лога
type Log struct {
	Path    string
	Params  map[string]string
	Message string
	Level   logLevel
	Time    time.Time
	TaskID  *string
}

type uuidKeyType string

const uuidKey uuidKeyType = "uuid"

// loggerSettings - Конфигурация логгера
type loggerSettings struct {
	isOn bool
}

var logger = &loggerSettings{
	isOn: true,
}

// Off выключает логгер
func Off() {
	logger.isOn = false
}

// Init конфигурирует логгер
func Init() {}

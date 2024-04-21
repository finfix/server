package logging

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

// Logger - Структура общего логгера, чтобы можно было легко заменить его
type Logger struct {
	*loggerSettings
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

// GetLogger возвращает логгер из любого места программы
func GetLogger() *Logger {
	return &Logger{logger}
}

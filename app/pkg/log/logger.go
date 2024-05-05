package log

import "server/app/pkg/errors"

// Log - Структура лога
type Log struct {
	Level   logLevel          `json:"level"`
	Message string            `json:"message"`
	Path    []string          `json:"path"`
	Params  map[string]string `json:"params,omitempty"`
	TaskID  *string           `json:"taskID,omitempty"`
}

type LogFormat string

const (
	JSONFormat LogFormat = "json"
	TextFormat LogFormat = "text"
)

func validateLogFormat(format LogFormat) error {
	switch format {
	case JSONFormat, TextFormat:
		return nil
	default:
		return errors.InternalServer.New("invalid log format")
	}
}

type uuidKeyType string

const uuidKey uuidKeyType = "uuid"

// loggerSettings - Конфигурация логгера
type loggerSettings struct {
	isOn      bool
	logFormat LogFormat
}

var logger = &loggerSettings{
	isOn:      true,
	logFormat: JSONFormat,
}

// Off выключает логгер
func Off() {
	logger.isOn = false
}

// Init конфигурирует логгер
func Init(logFormat LogFormat) error {

	if err := validateLogFormat(logFormat); err != nil {
		return err
	}
	logger.logFormat = logFormat

	return nil
}

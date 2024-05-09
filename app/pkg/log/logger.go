package log

import "server/app/pkg/errors"

// Log - Структура лога
type Log struct {
	Level            logLevel          `json:"level"`
	Message          string            `json:"message"`
	Path             []string          `json:"path"`
	Params           map[string]string `json:"params,omitempty"`
	AdditionalFields map[string]string `json:"additionalFields,omitempty"`
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

// loggerSettings - Конфигурация логгера
type loggerSettings struct {
	isOn             bool
	logFormat        LogFormat
	additionalFields map[string]string
}

var logger = &loggerSettings{
	isOn:             true,
	logFormat:        JSONFormat,
	additionalFields: make(map[string]string),
}

// Off выключает логгер
func Off() {
	logger.isOn = false
}

// Init конфигурирует логгер
func Init(
	logFormat LogFormat,
	additionalFields map[string]string,
) error {

	if err := validateLogFormat(logFormat); err != nil {
		return err
	}
	logger.logFormat = logFormat
	logger.additionalFields = additionalFields

	return nil
}

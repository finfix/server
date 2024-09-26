package log

import (
	"context"
	"os"
	"time"

	"pkg/log/model"
)

// loggerSettings - конфигурация логгера
type loggerSettings struct {

	// Массив обработчиков лога
	handlers []Handler

	// Дополнительные параметры, которые добавляются в каждый тег и настраиваются при инициализации
	systemInfo model.SystemInfo
}

// logger - синглтон переменная логгера
var logger = &loggerSettings{
	handlers: []Handler{
		NewConsoleHandler(os.Stdout, LevelDebug),
	},
	systemInfo: model.SystemInfo{
		Hostname: "",
		Version:  "",
		Build:    "",
		Env:      "",
	},
}

// Init конфигурирует логгер
func Init(
	systemInfo model.SystemInfo,
	handlers ...Handler,
) {
	logger = &loggerSettings{
		systemInfo: systemInfo,
		handlers:   handlers,
	}
}

func Off() {
	logger = new(loggerSettings)
}

// Error логгирует сообщения для ошибок системы
func Error(ctx context.Context, log any, opts ...Option) {
	for _, handler := range logger.handlers {
		handler.handle(ctx, LevelError, log, opts...)
	}
}

// Warning логгирует сообщения для ошибок пользователя
func Warning(ctx context.Context, log any, opts ...Option) {
	for _, handler := range logger.handlers {
		handler.handle(ctx, LevelWarning, log, opts...)
	}
}

// Info логгирует сообщения для информации
func Info(ctx context.Context, log any, opts ...Option) {
	for _, handler := range logger.handlers {
		handler.handle(ctx, LevelInfo, log, opts...)
	}
}

// Fatal логгирует сообщения для фатальных ошибок и завершает работу программы
func Fatal(ctx context.Context, log any, opts ...Option) {
	for _, handler := range logger.handlers {
		handler.handle(ctx, LevelFatal, log, opts...)
	}
	time.Sleep(1 * time.Second)
	os.Exit(1)
}

// Debug логгирует сообщения для дебага
func Debug(ctx context.Context, log any, opts ...Option) {
	for _, handler := range logger.handlers {
		handler.handle(ctx, LevelDebug, log, opts...)
	}
}

func GetSystemInfo() model.SystemInfo {
	return logger.systemInfo
}

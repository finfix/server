package logging

import (
	"os"
	"time"
)

// Panic логгирует сообщения при панике
func (logger *Logger) Panic(err error) {
	processingErrorLog(panicLevel, err)
}

// Error логгирует сообщения для ошибок системы
func (logger *Logger) Error(err error) {
	processingErrorLog(errorLevel, err)
}

// Warning логгирует сообщения для ошибок пользователя
func (logger *Logger) Warning(err error) {
	processingErrorLog(warningLevel, err)
}

// Info логгирует сообщения для информации
func (logger *Logger) Info(msg string, args ...any) {
	processingLog(infoLevel, msg, args...)
}

// Fatal логгирует сообщения для фатальных ошибок
func (logger *Logger) Fatal(err error) {
	processingErrorLog(fatalLevel, err)
	time.Sleep(1 * time.Second)
	os.Exit(1)
}

// Debug логгирует сообщения для дебага
func (logger *Logger) Debug(msg string, args ...any) {
	processingLog(debugLevel, msg, args...)
}

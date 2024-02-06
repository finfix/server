package logging

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"logger/app/logging/enum"
	pb "logger/app/pblogger"
	"pkg/errors"
)

// loggerSettings - Конфигурация логгера
type loggerSettings struct {
	loggingServer pb.LoggerClient
	serviceName   string
	isOff         bool
}

// Logger - Структура общего логгера, чтобы можно было легко заменить его
type Logger struct {
	*loggerSettings
}

var s = &loggerSettings{}

func Off() {
	s.isOff = true
}

// Init конфигурирует логгер
func Init(loggingServer pb.LoggerClient, serviceName string) {
	s = &loggerSettings{
		loggingServer: loggingServer,
		serviceName:   serviceName,
	}
}

// GetLogger возвращает логгер из любого места программы
func GetLogger() *Logger {
	return &Logger{s}
}

// DefaultErrorLoggerFunc логгирует ошибки по умолчанию
func DefaultErrorLoggerFunc(err error) {

	errorType := errors.GetType(err)
	logger := GetLogger()

	if errorType == errors.NoType {
		err = errors.InternalServer.WrapCtx(err, "Ошибка не имеет обертки, путь неверный")
		errorType = errors.InternalServer
	}

	switch errorType {
	case errors.InternalServer, errors.Forbidden, errors.BadGateway:
		logger.Error(err)
	default:
		logger.Warning(err)
	}
}

func DefaultRequestLoggerFunc(r *http.Request) {
	// Логгируем сообщение по типу METHOD /path
	GetLogger().Info("%s %s", r.Method, r.URL.Path)
}

func (logger *Logger) Panic(err error) {
	processingErrorLog(enum.Panic, true, err)
}

// Error логгирует сообщения для ошибок системы
func (logger *Logger) Error(err error) {
	processingErrorLog(enum.Error, true, err)
}

// Warning логгирует сообщения для ошибок пользователя
func (logger *Logger) Warning(err error) {
	processingErrorLog(enum.Warning, true, err)
}

// Info логгирует сообщения для информации
func (logger *Logger) Info(msg string, args ...any) {
	processingLog(enum.Info, msg, args...)
}

// Fatal логгирует сообщения для фатальных ошибок
func (logger *Logger) Fatal(err error) {
	processingErrorLog(enum.Fatal, true, err)
	time.Sleep(1 * time.Second)
	os.Exit(1)
}

// Debug логгирует сообщения для дебага
func (logger *Logger) Debug(msg string, args ...any) {
	processingLog(enum.Debug, msg, args...)
}

// processingErrorLog обрабатывает ошибки для логгирования
func processingErrorLog(level enum.LogLevel, sendToLogger bool, err error) {

	// Приводим пришедшую ошибку к нашей кастомной ошибке
	customErr, ok := err.(errors.CustomError)
	if !ok {
		context := "Ошибка приведения типов для CustomErr"
		customErr = errors.CustomError{
			DevelopText: err.Error(),
			Context:     &context,
		}
	}

	// Получаем текст первоначальной ошибки
	customErr.DevelopText = customErr.Error()

	// Переносим данные от ошибки к структуре лога
	var values *pb.Log = &pb.Log{}
	values.Path = customErr.Path
	values.Context = customErr.Context
	values.Message = customErr.DevelopText
	values.Level = string(level)
	values.Time = time.Now().Format("2006-01-02 15:04:05")
	values.Service = s.serviceName

	shareLog(values, sendToLogger)
}

// processingLog обрабатывает входные данные для логгирования
func processingLog(level enum.LogLevel, msg string, args ...any) {

	_, file, line, _ := runtime.Caller(2)

	// Переносим полученные данные к структуре лога
	var values *pb.Log = &pb.Log{}
	values.Path = file + ":" + strconv.Itoa(line)
	values.Message = fmt.Sprintf(msg, args...)
	values.Level = string(level)
	values.Time = time.Now().Format("2006-01-02 15:04:05")
	values.Service = s.serviceName

	shareLog(values, true)
}

func shareLog(values *pb.Log, isSendToLogger bool) {

	if s.isOff {
		return
	}

	// Выводим лог в консоль
	consoleLog := getConsoleLog(values)
	fmt.Println(consoleLog)

	// Отправляем лог на сервер
	if isSendToLogger && s.loggingServer != nil {
		go func() {
			if _, err := s.loggingServer.AddLog(context.Background(), &pb.Log{
				Level:   string(values.Level),
				Path:    values.Path,
				Context: values.Context,
				Message: values.Message,
				Service: values.Service,
				Time:    values.Time,
			}); err != nil {
				processingErrorLog(enum.Error, false, errors.InternalServer.WrapCtx(err, "Не смогли отправить лог на сервер"))
			}
		}()
	}
}

// getConsoleLog возвращает цветной лог из входных данных
func getConsoleLog(values *pb.Log) (log string) {

	// Окрашиваем заголовок лога
	switch enum.LogLevel(values.Level) {
	case enum.Error:
		log = "\x1b[31m[ERROR] \x1b[0m"
	case enum.Fatal:
		log = "\x1b[35m[FATAL] \x1b[0m"
	case enum.Info:
		log = "\x1b[34m[INFO] \x1b[0m"
	case enum.Debug:
		log = "\x1b[36m[DEBUG] \x1b[0m"
	case enum.Warning:
		log = "\x1b[33m[WARN] \x1b[0m"
	case enum.Panic:
		log = "\x1b[32m[PANIC] \x1b[0m"
	}

	log += fmt.Sprintf("%s %s", values.Path, values.Message)

	if values.Context != nil {
		log += fmt.Sprintf(". context: %s", *values.Context)
	}

	return log
}

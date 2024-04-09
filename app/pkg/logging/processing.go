package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"server/app/pkg/errors"
)

const (
	spacer = " "
)

var logHeaders = map[logLevel]string{
	errorLevel:   "\x1b[31m[ERROR] \x1b[0m ",
	fatalLevel:   "\x1b[35m[FATAL] \x1b[0m ",
	infoLevel:    "\x1b[34m[INFO] \x1b[0m ",
	debugLevel:   "\x1b[36m[DEBUG] \x1b[0m ",
	warningLevel: "\x1b[33m[WARN] \x1b[0m ",
	panicLevel:   "\x1b[32m[PANIC] \x1b[0m ",
}

// processingErrorLog обрабатывает ошибки для логгирования
func processingErrorLog(level logLevel, err error) {

	// Приводим пришедшую ошибку к нашей кастомной ошибке
	customErr := errors.CastError(err)

	shareLog(Log{
		Path:    customErr.Path,
		Params:  customErr.Params,
		Message: customErr.Error(),
		Level:   level,
		Time:    time.Now(),
	})
}

// processingLog обрабатывает входные данные для логгирования
func processingLog(level logLevel, msg string, args ...any) {

	_, file, line, _ := runtime.Caller(errors.SecondPathDepth)

	shareLog(Log{
		Path:    file + ":" + strconv.Itoa(line),
		Params:  nil,
		Message: fmt.Sprintf(msg, args...),
		Level:   level,
		Time:    time.Now(),
	})
}

func shareLog(values Log) {

	if logger.isOff {
		return
	}

	// Выводим лог в консоль
	consoleLog := getConsoleLog(values)
	fmt.Println(consoleLog) //nolint:forbidigo
}

// getConsoleLog возвращает цветной лог из входных данных
func getConsoleLog(values Log) (log string) {

	buffer := bytes.Buffer{}
	buffer.WriteString(logHeaders[values.Level])
	buffer.WriteString(values.Path)
	buffer.WriteString(spacer)
	buffer.WriteString(values.Message)
	buffer.WriteString(spacer)

	if len(values.Params) != 0 {
		params, err := json.Marshal(values.Params)
		if err != nil {
			GetLogger().Error(errors.InternalServer.Wrap(err))
		}
		buffer.WriteString(string(params))
	}

	return buffer.String()
}

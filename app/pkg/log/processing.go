package log

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"server/app/pkg/errors"
)

const (
	spacer = " "
)

var logHeaders = map[logLevel]string{
	errorLevel:   "\x1b[31mERROR\x1b[0m",
	fatalLevel:   "\x1b[35mFATAL\x1b[0m",
	infoLevel:    "\x1b[34mINFO\x1b[0m",
	debugLevel:   "\x1b[36mDEBUG\x1b[0m",
	warningLevel: "\x1b[33mWARN\x1b[0m",
	panicLevel:   "\x1b[32mPANIC\x1b[0m",
}

// processingErrorLog обрабатывает ошибки для логгирования
func processingErrorLog(ctx context.Context, level logLevel, err error) {

	// Приводим пришедшую ошибку к нашей кастомной ошибке
	customErr := errors.CastError(err)

	shareLog(Log{
		Path:    customErr.Path,
		Params:  customErr.Params,
		Message: customErr.Error(),
		Level:   level,
		Time:    time.Now(),
		TaskID:  ExtractTaskID(ctx),
	})
}

// processingLog обрабатывает входные данные для логгирования
func processingLog(ctx context.Context, level logLevel, msg string, args ...any) {

	_, file, line, _ := runtime.Caller(errors.SecondPathDepth)

	shareLog(Log{
		Path:    file + ":" + strconv.Itoa(line),
		Params:  nil,
		Message: fmt.Sprintf(msg, args...),
		Level:   level,
		Time:    time.Now(),
		TaskID:  ExtractTaskID(ctx),
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

	logComponents := []string{
		logHeaders[values.Level],
		values.Time.Format("2006-01-02 15:04:05.000"),
	}

	if values.TaskID != nil {
		logComponents = append(logComponents, *values.TaskID)
	}

	logComponents = append(logComponents, values.Path, values.Message)

	if len(values.Params) != 0 {
		params, err := json.Marshal(values.Params)
		if err != nil {
			Error(context.Background(), errors.InternalServer.Wrap(err))
		}
		logComponents = append(logComponents, string(params))
	}

	return strings.Join(logComponents, spacer)
}

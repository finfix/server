package log

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"server/app/pkg/errors"
	"server/app/pkg/maps"
)

const (
	spacer = " "
)

var logHeaders = map[logLevel]string{
	errorLevel:   "\x1b[31mERROR\x1b[0m",
	fatalLevel:   "\x1b[35mFATAL\x1b[0m",
	infoLevel:    "\x1b[34mINFO \x1b[0m",
	debugLevel:   "\x1b[36mDEBUG\x1b[0m",
	warningLevel: "\x1b[33mWARN \x1b[0m",
	panicLevel:   "\x1b[32mPANIC\x1b[0m",
}

// processingErrorLog обрабатывает ошибки для логгирования
func processingErrorLog(ctx context.Context, level logLevel, err error) {

	// Приводим пришедшую ошибку к нашей кастомной ошибке
	customErr := errors.CastError(err)

	shareLog(Log{
		Path:             customErr.Path,
		Params:           customErr.Params,
		Message:          customErr.Error(),
		Level:            level,
		AdditionalFields: maps.Join(logger.additionalFields, ExtractAdditionalInfo(ctx)),
	})
}

// processingLog обрабатывает входные данные для логгирования
func processingLog(ctx context.Context, level logLevel, msg string, opts ...Option) {

	options := mergeOptions(opts...)

	shareLog(Log{
		Path:             errors.GetPath(errors.ThirdPathDepth),
		Params:           options.params,
		Message:          msg,
		Level:            level,
		AdditionalFields: maps.Join(logger.additionalFields, ExtractAdditionalInfo(ctx)),
	})
}

func shareLog(values Log) {

	// Если логи выключены, то не пишем их
	if !logger.isOn {
		return
	}

	// Определяем формат лога и получаем строку лога
	var logLine string
	switch logger.logFormat {
	case TextFormat:
		logLine = getConsoleLog(values)
	case JSONFormat:
		logLine = getJSONLog(values)
	}

	// Определяем в какой поток писать лог
	var writer io.Writer
	switch values.Level {
	case errorLevel, fatalLevel, warningLevel, panicLevel:
		writer = os.Stderr
	case infoLevel, debugLevel:
		writer = os.Stdout
	}

	// Пишем лог в консоль в выбранный поток
	if _, err := io.WriteString(writer, logLine+"\n"); err != nil {
		log.Println(err)
	}
}

// getConsoleLog возвращает цветной лог из входных данных
func getConsoleLog(values Log) string {

	logComponents := []string{
		logHeaders[values.Level], // Цветной заголовок с уровнем лога
		values.Path[0],           // Путь к месту, где был вызван лог (или где была создана ошибка)
		values.Message,           // Сообщение лога
	}

	return strings.Join(logComponents, spacer)
}

// getJSONLog возвращает JSONFormat лог из входных данных
func getJSONLog(values Log) string {

	log, err := json.Marshal(values)
	if err != nil {
		Error(context.Background(), errors.InternalServer.Wrap(err))
	}

	return string(log)
}

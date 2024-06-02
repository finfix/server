package log

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/hasher"
)

func handleProcessing(ctx context.Context, level logLevel, log any, opts ...Option) {
	switch log := log.(type) {
	case string:
		processingStringLog(ctx, level, log, opts...)
	case error:
		processingErrorLog(ctx, level, log)
	default:
		processingErrorLog(ctx, errorLevel, errors.InternalServer.New(
			fmt.Sprintf("Processor log for type %T not implemented", log),
			[]errors.Option{
				errors.PathDepthOption(errors.ThirdPathDepth),
			}...,
		))
	}
}

// Error логгирует сообщения для ошибок системы
func Error(ctx context.Context, log any) {
	handleProcessing(ctx, errorLevel, log)
}

// Warning логгирует сообщения для ошибок пользователя
func Warning(ctx context.Context, log any) {
	handleProcessing(ctx, warningLevel, log)
}

// Info логгирует сообщения для информации
func Info(ctx context.Context, log any, opts ...Option) {
	handleProcessing(ctx, infoLevel, log)
}

// Fatal логгирует сообщения для фатальных ошибок
func Fatal(ctx context.Context, log error) {
	handleProcessing(ctx, fatalLevel, log)
	time.Sleep(1 * time.Second)
	os.Exit(1)
}

// Debug логгирует сообщения для дебага
func Debug(ctx context.Context, log string, opts ...Option) {
	handleProcessing(ctx, debugLevel, log)
}

// SetTaskID устанавливает TaskID в контекст
func SetTaskID(ctx context.Context) context.Context {
	uuid, err := generateTaskID()
	if err != nil {
		Error(ctx, err)
		return ctx
	}
	return contextKeys.SetTaskID(ctx, uuid)
}

func generateTaskID() (string, error) {
	var length uint32 = 4
	bytes, err := hasher.GenerateRandomBytes(length)
	if err != nil {
		return "", errors.InternalServer.Wrap(err)
	}
	return fmt.Sprintf("%x", bytes[:4]), nil
}

// ExtractAdditionalInfo извлекает дополнительную информацию из контекста
func ExtractAdditionalInfo(ctx context.Context) map[string]string {

	additionalInfo := make(map[string]string)

	if ctx == nil {
		return additionalInfo
	}

	if taskID := contextKeys.GetTaskID(ctx); taskID != nil {
		additionalInfo["taskID"] = *taskID
	}

	if userID := contextKeys.GetUserID(ctx); userID != nil && *userID != 0 {
		additionalInfo["userID"] = strconv.Itoa(int(*userID))
	}

	if deviceID := contextKeys.GetDeviceID(ctx); deviceID != nil {
		additionalInfo["deviceID"] = *deviceID
	}

	return additionalInfo
}

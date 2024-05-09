package log

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"time"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
)

// Error логгирует сообщения для ошибок системы
func Error(ctx context.Context, err error) {
	processingErrorLog(ctx, errorLevel, err)
}

// Warning логгирует сообщения для ошибок пользователя
func Warning(ctx context.Context, err error) {
	processingErrorLog(ctx, warningLevel, err)
}

// Info логгирует сообщения для информации
func Info(ctx context.Context, msg string, opts ...Option) {
	processingLog(ctx, infoLevel, msg, opts...)
}

// Fatal логгирует сообщения для фатальных ошибок
func Fatal(ctx context.Context, err error) {
	processingErrorLog(context.Background(), fatalLevel, err)
	time.Sleep(1 * time.Second)
	os.Exit(1)
}

// Debug логгирует сообщения для дебага
func Debug(ctx context.Context, msg string, opts ...Option) {
	processingLog(ctx, debugLevel, msg, opts...)
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
	length := 4
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", errors.InternalServer.Wrap(err)
	}
	return fmt.Sprintf("%x", b[:4]), nil
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

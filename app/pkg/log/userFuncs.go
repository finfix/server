package log

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"server/app/pkg/errors"
)

// Panic логгирует сообщения при панике
func Panic(ctx context.Context, err error) {
	processingErrorLog(ctx, panicLevel, err)
}

// Error логгирует сообщения для ошибок системы
func Error(ctx context.Context, err error) {
	processingErrorLog(ctx, errorLevel, err)
}

// Warning логгирует сообщения для ошибок пользователя
func Warning(ctx context.Context, err error) {
	processingErrorLog(ctx, warningLevel, err)
}

// Info логгирует сообщения для информации
func Info(ctx context.Context, msg string, args ...any) {
	processingLog(ctx, infoLevel, msg, args...)
}

// Fatal логгирует сообщения для фатальных ошибок
func Fatal(err error) {
	processingErrorLog(context.Background(), fatalLevel, err)
	time.Sleep(1 * time.Second)
	os.Exit(1)
}

// Debug логгирует сообщения для дебага
func Debug(ctx context.Context, msg string, args ...any) {
	processingLog(ctx, debugLevel, msg, args...)
}

// SetTaskID устанавливает TaskID в контекст
func SetTaskID(ctx context.Context) context.Context {
	uuid, err := generateTaskID()
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, uuidKey, uuid)
}

func generateTaskID() (string, error) {
	length := 4
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", errors.InternalServer.Wrap(err)
	}
	return fmt.Sprintf("%x", b[:4]), nil
}

// ExtractTaskID извлекает TaskID из контекста
func ExtractTaskID(ctx context.Context) *string {
	if ctx == nil {
		return nil
	}
	if uuid, ok := ctx.Value(uuidKey).(string); ok {
		return &uuid
	}
	return nil
}

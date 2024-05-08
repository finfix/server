package contextKeys

import "context"

type contextKey int

const (
	deviceIDKey contextKey = iota + 1
	userIDKey
	taskIDKey
)

func SetDeviceID(ctx context.Context, deviceID string) context.Context {
	return context.WithValue(ctx, deviceIDKey, deviceID)
}

func SetUserID(ctx context.Context, userID uint32) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func SetTaskID(ctx context.Context, taskID string) context.Context {
	return context.WithValue(ctx, taskIDKey, taskID)
}

func GetDeviceID(ctx context.Context) *string {
	if v, ok := ctx.Value(deviceIDKey).(string); ok {
		return &v
	}
	return nil
}

func GetUserID(ctx context.Context) *uint32 {
	if v, ok := ctx.Value(userIDKey).(uint32); ok {
		return &v
	}
	return nil
}

func GetTaskID(ctx context.Context) *string {
	if v, ok := ctx.Value(taskIDKey).(string); ok {
		return &v
	}
	return nil
}

package contextKeys

import "context"

type ContextKey int

const (
	DeviceIDKey ContextKey = iota + 1
	UserIDKey
	TaskIDKey
)

func SetDeviceID(ctx context.Context, deviceID string) context.Context {
	return context.WithValue(ctx, DeviceIDKey, deviceID)
}

func SetUserID(ctx context.Context, userID uint32) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

func SetTaskID(ctx context.Context, taskID string) context.Context {
	return context.WithValue(ctx, TaskIDKey, taskID)
}

func GetDeviceID(ctx context.Context) *string {
	if v, ok := ctx.Value(DeviceIDKey).(string); ok {
		return &v
	}
	return nil
}

func GetUserID(ctx context.Context) *uint32 {
	if v, ok := ctx.Value(UserIDKey).(uint32); ok {
		return &v
	}
	return nil
}

func GetTaskID(ctx context.Context) *string {
	if v, ok := ctx.Value(TaskIDKey).(string); ok {
		return &v
	}
	return nil
}

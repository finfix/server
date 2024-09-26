package contextKeys

import (
	"context"

	"pkg/log/model"
)

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

func SetRequestID(ctx context.Context, taskID string) context.Context {
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

func GetRequestID(ctx context.Context) *string {
	if v, ok := ctx.Value(TaskIDKey).(string); ok {
		return &v
	}
	return nil
}

// GetUserInfo извлекает дополнительную информацию из контекста
func GetUserInfo(ctx context.Context) *model.UserInfo {

	var userInfo model.UserInfo

	if ctx == nil {
		return nil
	}

	if userID := GetUserID(ctx); userID != nil && *userID != 0 {
		userInfo.UserID = userID
	}
	userInfo.TaskID = GetRequestID(ctx)
	userInfo.DeviceID = GetDeviceID(ctx)

	return &userInfo
}

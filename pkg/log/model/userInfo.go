package model

type UserInfo struct {
	UserID   *uint32 `json:"userID"`
	TaskID   *string `json:"taskID"`
	DeviceID *string `json:"deviceID"`
}

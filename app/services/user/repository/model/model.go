package model

import (
	"server/app/services/user/model/OS"
)

type UpdateDeviceReq struct {
	UserID            uint32
	DeviceID          string
	RefreshToken      *string
	NotificationToken *string
}

type CreateDeviceReq struct {
	RefreshToken string
	DeviceID     string
	UserID       uint32
	OS           OS.OS
	BundleID     string
}

type GetDevicesReq struct {
	IDs       []uint32
	DeviceIDs []string
	UserIDs   []uint32
}

type UpdateUserReq struct {
	ID              uint32
	Name            *string
	Email           *string
	PasswordHash    *[]byte
	PasswordSalt    *[]byte
	DefaultCurrency *string
}

package model

type UpdateDeviceReq struct {
	UserID                 uint32
	DeviceID               string
	RefreshToken           *string
	NotificationToken      *string
	ApplicationInformation UpdateApplicationInformationReq
	DeviceInformation      UpdateDeviceInformationReq
}

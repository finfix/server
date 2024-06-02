package model

type UpdateDeviceReq struct {
	UserID                 uint32
	DeviceID               string
	RefreshToken           *string
	NotificationToken      *string
	ApplicationInformation UpdateApplicationInformationReq
	DeviceInformation      UpdateDeviceInformationReq
}

type UpdateApplicationInformationReq struct {
	BundleID *string
	Version  *string
	Build    *string
}

type UpdateDeviceInformationReq struct {
	VersionOS *string
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

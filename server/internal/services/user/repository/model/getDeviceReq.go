package model

type GetDevicesReq struct {
	IDs       []uint32
	DeviceIDs []string
	UserIDs   []uint32
}

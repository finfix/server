package testingFunc

const (
	InvalidJSON = "{invalid}"
)

var (
	GeneralCtx = NewCtxBuilder(make(map[string]any)).
		Set("DeviceID", "DeviceID").
		Set("UserID", uint32(1))
)

package constants

type CommonFields struct {
	// TODO: Убрать json теги
	UserID   uint32 `json:"userID" schema:"-" validate:"required" minimum:"1"` // Идентификатор пользователя
	DeviceID string `json:"deviceID" schema:"-" validate:"required"`           // Идентификатор устройства
}

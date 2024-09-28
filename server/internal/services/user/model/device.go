package model

type Device struct {
	DeviceInformation              // Информация о девайсе пользователя
	ApplicationInformation         // Информация о приложении пользователя
	NotificationToken      *string `db:"notification_token" json:"-"` // Токен для системы уведомлений
	RefreshToken           string  `db:"refresh_token" json:"-"` // Токен доступа для обновления пары токенов
	UserID                 uint32  `db:"user_id" json:"userID"` // Идентификатор пользователя девайса
	DeviceID               string  `db:"device_id" json:"deviceID"` // Идентификатор девайса
}

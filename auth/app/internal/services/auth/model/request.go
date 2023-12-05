package model

type RefreshTokensReq struct {
	Token string // Токен восстановления доступа
}

type SignInReq struct {
	Email    string // Электронная почта пользователя
	Password string // Пароль пользователя
	DeviceID string // Идентификатор устройства
}

type SignUpReq struct {
	Name     string // Имя пользователя
	Email    string // Электронная почта пользователя
	Password string // Пароль пользователя
	DeviceID string // Идентификатор устройства
}

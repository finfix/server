package model

type RefreshTokensReq struct {
	Token string `json:"token" validate:"required"` // Токен восстановления доступа
}

type SignInReq struct {
	Email    string `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password string `json:"password" validate:"required"`             // Пароль пользователя
	DeviceID string `json:"-" validate:"required"`                    // Идентификатор устройства
}

type SignUpReq struct {
	Name     string `json:"name" validate:"required"`                 // Имя пользователя
	Email    string `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password string `json:"password" validate:"required"`             // Пароль пользователя
	DeviceID string `json:"-" validate:"required"`                    // Идентификатор устройства
}

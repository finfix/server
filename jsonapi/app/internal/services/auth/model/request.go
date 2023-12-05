package model

type RefreshTokensReq struct {
	Token string `jsonapi:"token" validate:"required"` // Токен восстановления доступа
}

type SignInReq struct {
	Email    string `jsonapi:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password string `jsonapi:"password" validate:"required"`             // Пароль пользователя
	DeviceID string `jsonapi:"-" validate:"required"`                    // Идентификатор устройства
}

type SignUpReq struct {
	Name     string `jsonapi:"name" validate:"required"`                 // Имя пользователя
	Email    string `jsonapi:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password string `jsonapi:"password" validate:"required"`             // Пароль пользователя
	DeviceID string `jsonapi:"-" validate:"required"`                    // Идентификатор устройства
}

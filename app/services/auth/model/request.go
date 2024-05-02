package model

import (
	"server/app/pkg/datetime"
	"server/app/services"
)

type RefreshTokensReq struct {
	Token     string `json:"token" validate:"required"` // Токен восстановления доступа
	Necessary services.NecessaryUserInformation
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

type Session struct {
	ExpiresAt datetime.Time `db:"expires_at"`
	ID        uint32        `db:"id"`
	DeviceID  string        `db:"device_id"`
}

package model

import (
	"server/app/pkg/datetime"
	"server/app/pkg/validation"
	"server/app/services"
)

type RefreshTokensReq struct {
	Token     string `json:"token" validate:"required"` // Токен восстановления доступа
	Necessary services.NecessaryUserInformation
}

func (r RefreshTokensReq) Validate() error { return nil }

func (r RefreshTokensReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	r.Necessary = necessary
	return r
}

type SignInReq struct {
	Email    string `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password string `json:"password" validate:"required"`             // Пароль пользователя
	DeviceID string `json:"-" validate:"required"`                    // Идентификатор устройства
}

func (r SignInReq) Validate() error {
	return validation.Mail(r.Email)
}

func (r SignInReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	r.DeviceID = necessary.DeviceID
	return r
}

type SignUpReq struct {
	Name     string `json:"name" validate:"required"`                 // Имя пользователя
	Email    string `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password string `json:"password" validate:"required"`             // Пароль пользователя
	DeviceID string `json:"-" validate:"required"`                    // Идентификатор устройства
}

func (r SignUpReq) Validate() error {
	return validation.Mail(r.Email)
}

func (r SignUpReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	r.DeviceID = necessary.DeviceID
	return r
}

type Session struct {
	ExpiresAt datetime.Time `db:"expires_at"`
	ID        uint32        `db:"id"`
	DeviceID  string        `db:"device_id"`
}

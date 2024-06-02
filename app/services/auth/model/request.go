package model

import (
	"server/app/pkg/validation"
	"server/app/services"
	userModel "server/app/services/user/model"
)

type RefreshTokensReq struct {
	Token       string                           `json:"token" validate:"required"` // Токен восстановления доступа
	Application userModel.ApplicationInformation `json:"application"`               // Информация о приложении
	Device      userModel.DeviceInformation      `json:"device"`                    // Информация о девайсе
	Necessary   services.NecessaryUserInformation
}

func (r RefreshTokensReq) Validate() error { return nil }

func (r RefreshTokensReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	r.Necessary = necessary
	return r
}

type SignInReq struct {
	Email       string                           `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password    string                           `json:"password" validate:"required"`             // Пароль пользователя
	Application userModel.ApplicationInformation `json:"application"`                              // Информация о приложении
	Device      userModel.DeviceInformation      `json:"device"`                                   // Информация о девайсе
	DeviceID    string                           `json:"-" validate:"required"`                    // Идентификатор устройства
}

func (r SignInReq) Validate() error {
	return validation.Mail(r.Email)
}

func (r SignInReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	r.DeviceID = necessary.DeviceID
	return r
}

type SignUpReq struct {
	Name        string                           `json:"name" validate:"required"`                 // Имя пользователя
	Email       string                           `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password    string                           `json:"password" validate:"required"`             // Пароль пользователя
	Application userModel.ApplicationInformation `json:"application"`                              // Информация о приложении
	Device      userModel.DeviceInformation      `json:"device"`                                   // Информация о девайсе
	DeviceID    string                           `json:"-" validate:"required"`                    // Идентификатор устройства
}

func (r SignUpReq) Validate() error {
	return validation.Mail(r.Email)
}

func (r SignUpReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	r.DeviceID = necessary.DeviceID
	return r
}

type SignOutReq struct {
	Necessary services.NecessaryUserInformation
}

func (r SignOutReq) Validate() error {
	return nil
}

func (r SignOutReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	r.Necessary = necessary
	return r
}

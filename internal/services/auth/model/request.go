package model

import (
	"server/internal/services"
	userModel "server/internal/services/user/model"
)

type RefreshTokensReq struct {
	Token       string                            `json:"token" validate:"required"` // Токен восстановления доступа
	Application userModel.ApplicationInformation  `json:"application"`               // Информация о приложении
	Device      userModel.DeviceInformation       `json:"device"`                    // Информация о девайсе
	Necessary   services.NecessaryUserInformation `json:"-"`
}

type SignInReq struct {
	Email       string                           `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password    string                           `json:"password" validate:"required"`             // Пароль пользователя
	Application userModel.ApplicationInformation `json:"application"`                              // Информация о приложении
	Device      userModel.DeviceInformation      `json:"device"`                                   // Информация о девайсе
	DeviceID    string                           `json:"-" validate:"required"`                    // Идентификатор устройства
}

type SignUpReq struct {
	Name        string                           `json:"name" validate:"required"`                 // Имя пользователя
	Email       string                           `json:"email" validate:"required" format:"email"` // Электронная почта пользователя
	Password    string                           `json:"password" validate:"required"`             // Пароль пользователя
	Application userModel.ApplicationInformation `json:"application"`                              // Информация о приложении
	Device      userModel.DeviceInformation      `json:"device"`                                   // Информация о девайсе
	DeviceID    string                           `json:"-" validate:"required"`                    // Идентификатор устройства
}

type SignOutReq struct {
	Necessary services.NecessaryUserInformation
}

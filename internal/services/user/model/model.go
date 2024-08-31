package model

import (
	"server/pkg/datetime"
	"server/internal/services/user/model/OS"
)

type User struct {
	ID              uint32        `db:"id" json:"id"`                                      // Идентификатор пользователя
	Name            string        `db:"name" json:"name"`                                  // Имя пользователя
	Email           string        `db:"email" json:"email"`                                // Электронная почта
	PasswordHash    []byte        `db:"password_hash" json:"-"`                            // Хэш пароля
	PasswordSalt    []byte        `db:"password_salt" json:"-"`                            // Соль пароля
	TimeCreate      datetime.Time `db:"time_create" json:"timeCreate"`                     // Дата и время создания аккаунта
	DefaultCurrency string        `db:"default_currency_signatura" json:"defaultCurrency"` // Валюта по умолчанию
	IsAdmin         bool          `db:"is_admin" json:"-"`                                 // Является ли пользователь администратором системы
}

type ApplicationInformation struct {
	BundleID string `json:"bundleID" validate:"required" db:"application_bundle_id"` // Бандл приложения
	Version  string `json:"version" validate:"required" db:"application_version"`    // Версия приложения
	Build    string `json:"build" validate:"required" db:"application_build"`        // Билд приложения
}

type DeviceInformation struct {
	NameOS     OS.OS  `json:"nameOS" validate:"required" db:"device_os_name"`       // Название операционной системы
	VersionOS  string `json:"versionOS" validate:"required" db:"device_os_version"` // Версия операционной системы
	DeviceName string `json:"deviceName" validate:"required" db:"device_name"`      // Название девайса
	ModelName  string `json:"modelName" validate:"required" db:"device_model_name"` // Название модели
	IPAddress  string `json:"-" db:"device_ip_address"`                             // IP-адрес
	UserAgent  string `json:"-" db:"device_user_agent"`                             // UserAgent
}

type Device struct {
	DeviceInformation              // Информация о девайсе пользователя
	ApplicationInformation         // Информация о приложении пользователя
	NotificationToken      *string `db:"notification_token" json:"-"` // Токен для системы уведомлений
	RefreshToken           string  `db:"refresh_token" json:"-"` // Токен доступа для обновления пары токенов
	UserID                 uint32  `db:"user_id" json:"userID"` // Идентификатор пользователя девайса
	DeviceID               string  `db:"device_id" json:"deviceID"` // Идентификатор девайса
}

type Notification struct {
	Title      string `json:"title"`      // Заголовок уведомления
	Subtitle   string `json:"subtitle"`   // Подзаголовок уведомления
	Message    string `json:"message"`    // Сообщение уведомления
	BadgeCount uint8  `json:"badgeCount"` // Индикатор какое значение показывать в бадже
}

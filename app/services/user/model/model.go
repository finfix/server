package model

import (
	"server/app/pkg/datetime"
	"server/app/services/user/model/OS"
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

type Device struct {
	OS                OS.OS   `db:"os"`                 // Операционная система
	NotificationToken *string `db:"notification_token"` // Токен для системы уведомлений
	RefreshToken      string  `db:"refresh_token"`      // Токен доступа для обновления пары токенов
	UserID            uint32  `db:"user_id"`            // Идентификатор пользователя девайса
	DeviceID          string  `db:"device_id"`          // Идентификатор девайса
	BundleID          string  `db:"bundle_id"`          // Бандл приложения
}

type Notification struct {
	Title      string `json:"title"`      // Заголовок уведомления
	Subtitle   string `json:"subtitle"`   // Подзаголовок уведомления
	Message    string `json:"message"`    // Сообщение уведомления
	BadgeCount uint8  `json:"badgeCount"` // Индикатор какое значение показывать в бадже
}

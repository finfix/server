package model

import "server/app/pkg/datetime"

type User struct {
	ID                    uint32        `db:"id" json:"id"`                                      // Идентификатор пользователя
	Name                  string        `db:"name" json:"name"`                                  // Имя пользователя
	Email                 string        `db:"email" json:"email"`                                // Электронная почта
	PasswordHash          string        `db:"password_hash" json:"-"`                            // Хэш пароля
	VerificationEmailCode *string       `db:"verification_email_code" json:"-"`                  // Временный код, приходящий на почту
	TimeCreate            datetime.Time `db:"time_create" json:"timeCreate"`                     // Дата и время создания аккаунта
	FCMToken              *string       `db:"fcm_token" json:"-"`                                // Токен уведомлений
	DefaultCurrency       string        `db:"default_currency_signatura" json:"defaultCurrency"` // Валюта по умолчанию
}

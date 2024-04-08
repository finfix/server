package model

import "time"

type Currency struct {
	Signatura string  `json:"isoCode" db:"signatura"` // Сигнатура валюты
	Name      string  `json:"name" db:"name"`         // Название валюты
	Symbol    string  `json:"symbol" db:"symbol"`     // Символ валюты
	Rate      float64 `json:"rate" db:"rate"`         // Курс валюты
}

type User struct {
	ID                    uint32    `db:"id" json:"id"`                                      // Идентификатор пользователя
	Name                  string    `db:"name" json:"name"`                                  // Имя пользователя
	Email                 string    `db:"email" json:"email"`                                // Электронная почта
	PasswordHash          string    `db:"password_hash" json:"-"`                            // Хэш пароля
	VerificationEmailCode *string   `db:"verification_email_code" json:"-"`                  // Временный код, приходящий на почту
	TimeCreate            time.Time `db:"time_create" json:"timeCreate"`                     // Дата и время создания аккаунта
	FCMToken              *string   `db:"fcm_token" json:"-"`                                // Токен уведомлений
	DefaultCurrency       string    `db:"default_currency_signatura" json:"defaultCurrency"` // Валюта по умолчанию
}

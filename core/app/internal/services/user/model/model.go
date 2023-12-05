package model

import "time"

type User struct {
	ID                    uint32    `db:"id"`                         // Идентификатор пользователя
	Name                  string    `db:"name"`                       // Имя пользователя
	Email                 string    `db:"email"`                      // Электронная почта
	PasswordHash          string    `db:"password_hash"`              // Хэш пароля
	VerificationEmailCode *string   `db:"verification_email_code"`    // Временный код, приходящий на почту
	TimeCreate            time.Time `db:"time_create"`                // Дата и время создания аккаунта
	FCMToken              *string   `db:"fcm_token"`                  // Токен уведомлений
	DefaultCurrency       string    `db:"default_currency_signatura"` // Валюта по умолчанию
}

type Currency struct {
	Signatura string  `db:"signatura"` // Сигнатура валюты
	Name      string  `db:"name"`      // Название валюты
	Symbol    string  `db:"symbol"`    // Символ валюты
	Rate      float64 `db:"rate"`      // Курс валюты
}

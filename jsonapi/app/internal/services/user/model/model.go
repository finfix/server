package model

import "time"

type Currency struct {
	Signatura string  `json:"isoCode"`
	Name      string  `json:"name"`
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
}

type User struct {
	ID              uint32     `json:"id"`              // Идентификатор пользователя
	Name            string     `json:"name"`            // Имя пользователя
	Email           string     `json:"email"`           // Электронная почта
	TimeCreate      *time.Time `json:"timeCreate"`      // Дата и время создания аккаунта
	DefaultCurrency string     `json:"defaultCurrency"` // Валюта по умолчанию
}

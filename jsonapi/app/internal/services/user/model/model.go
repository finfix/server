package model

import "time"

type Currency struct {
	Signatura string  `jsonapi:"isoCode"`
	Name      string  `jsonapi:"name"`
	Symbol    string  `jsonapi:"symbol"`
	Rate      float64 `jsonapi:"rate"`
}

type User struct {
	ID              uint32     `jsonapi:"id"`              // Идентификатор пользователя
	Name            string     `jsonapi:"name"`            // Имя пользователя
	Email           string     `jsonapi:"email"`           // Электронная почта
	TimeCreate      *time.Time `jsonapi:"timeCreate"`      // Дата и время создания аккаунта
	DefaultCurrency string     `jsonapi:"defaultCurrency"` // Валюта по умолчанию
}

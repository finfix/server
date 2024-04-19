package model

type Currency struct {
	Signatura string  `json:"isoCode" db:"signatura"` // Сигнатура валюты
	Name      string  `json:"name" db:"name"`         // Название валюты
	Symbol    string  `json:"symbol" db:"symbol"`     // Символ валюты
	Rate      float64 `json:"rate" db:"rate"`         // Курс валюты
}

type Version struct {
	Version string `json:"version"` // Версия сервера
	Build   string `json:"build"`   // Номер сборки
}

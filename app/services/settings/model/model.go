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

type Icon struct {
	ID   int    `json:"id" db:"id"`     // ID иконки
	Name string `json:"name" db:"name"` // Название иконки
	Url  string `json:"url" db:"img"`   // URL иконки
}

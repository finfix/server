package model

import "server/internal/services"

type UpdateAccountGroupReq struct {
	Necessary    services.NecessaryUserInformation
	ID           uint32  `json:"id" db:"id"`                       // Идентификатор группы счетов
	Name         *string `json:"name" db:"name"`                   // Название группы счетов
	Currency     *string `json:"currency" db:"currency_signatura"` // Валюта группы счетов
	Visible      *bool   `json:"visible" db:"visible"`             // Видимость группы счетов
	SerialNumber *uint32 `json:"serialNumber" db:"serial_number"`  // Порядковый номер группы счетов
}

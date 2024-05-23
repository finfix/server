package model

import "server/app/pkg/datetime"

type AccountGroup struct {
	ID             uint32        `json:"id" db:"id"`                          // Идентификатор группы счетов
	Name           string        `json:"name" db:"name"`                      // Название группы счетов
	Currency       string        `json:"currency" db:"currency_signatura"`    // Валюта группы счетов
	SerialNumber   uint32        `json:"serialNumber" db:"serial_number"`     // Порядковый номер группы счетов
	Visible        bool          `json:"visible" db:"visible"`                // Видимость группы счетов
	DatetimeCreate datetime.Time `json:"datetimeCreate" db:"datetime_create"` // Дата и время создания группы счетов
}

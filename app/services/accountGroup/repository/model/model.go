package model

import (
	"time"
)

type CreateAccountGroupReq struct {
	Name           string    // Название группы счетов
	Currency       string    // Валюта группы счетов
	Visible        bool      // Видимость группы счетов
	SerialNumber   uint32    // Порядковый номер группы счетов
	DatetimeCreate time.Time // Дата и время создания группы счетов
	UserID         uint32    // Каким пользователем создан объект
}

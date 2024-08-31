package model

import (
	"server/internal/services"
	repoModel "server/internal/services/accountGroup/repository/model"
	"server/pkg/datetime"
)

type CreateAccountGroupReq struct {
	Necessary      services.NecessaryUserInformation
	Name           string        `json:"name" db:"name" validate:"required"`                      // Название группы счетов
	Currency       string        `json:"currency" db:"currency_signatura" validate:"required"`    // Валюта группы счетов
	SerialNumber   uint32        `json:"serialNumber" db:"serial_number" validate:"required"`     // Порядковый номер группы счетов
	DatetimeCreate datetime.Time `json:"datetimeCreate" db:"datetime_create" validate:"required"` // Дата и время создания группы счетов
}

func (s CreateAccountGroupReq) ConvertToRepoReq() repoModel.CreateAccountGroupReq {
	return repoModel.CreateAccountGroupReq{
		UserID:         s.Necessary.UserID,
		Name:           s.Name,
		Currency:       s.Currency,
		Visible:        true,
		SerialNumber:   s.SerialNumber,
		DatetimeCreate: s.DatetimeCreate.Time,
	}
}

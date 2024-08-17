package model

import (
	"server/app/pkg/datetime"
	"server/app/services"
	repoModel "server/app/services/accountGroup/repository/model"
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

func (s CreateAccountGroupReq) Validate() error {
	return nil
}

type GetAccountGroupsReq struct {
	Necessary       services.NecessaryUserInformation
	AccountGroupIDs []uint32 `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"` // Идентификаторы групп счетов
}

type UpdateAccountGroupReq struct {
	Necessary    services.NecessaryUserInformation
	ID           uint32  `json:"id" db:"id"`                       // Идентификатор группы счетов
	Name         *string `json:"name" db:"name"`                   // Название группы счетов
	Currency     *string `json:"currency" db:"currency_signatura"` // Валюта группы счетов
	Visible      *bool   `json:"visible" db:"visible"`             // Видимость группы счетов
	SerialNumber *uint32 `json:"serialNumber" db:"serial_number"`  // Порядковый номер группы счетов
}

type DeleteAccountGroupReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" schema:"id" validate:"required" minimum:"1"` // Идентификатор счета
}

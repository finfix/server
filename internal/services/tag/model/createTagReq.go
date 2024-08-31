package model

import (
	"server/internal/services"
	repoModel "server/internal/services/tag/repository/model"
	"server/pkg/datetime"
)

type CreateTagReq struct {
	Necessary      services.NecessaryUserInformation
	AccountGroupID uint32        `json:"accountGroupID" validate:"required"` // Идентификатор группы счетов
	Name           string        `json:"name" validate:"required"`           // Название подкатегории
	DatetimeCreate datetime.Time `json:"datetimeCreate" validate:"required"` // Дата создания подкатегории
}

func (s CreateTagReq) ConvertToRepoReq() repoModel.CreateTagReq {
	return repoModel.CreateTagReq{
		Name:            s.Name,
		AccountGroupID:  s.AccountGroupID,
		CreatedByUserID: s.Necessary.UserID,
		DatetimeCreate:  s.DatetimeCreate.Time,
	}
}

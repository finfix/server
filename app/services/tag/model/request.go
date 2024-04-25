package model

import (
	"server/app/pkg/datetime"
	"server/app/services"
	repoModel "server/app/services/tag/repository/model"
)

type DeleteTagReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор подкатегории
}

type CreateTagReq struct {
	Necessary      services.NecessaryUserInformation
	AccountGroupID uint32        `json:"accountGroupID" validate:"required"` // Идентификатор группы счетов
	Name           string        `json:"name" validate:"required"`           // Название подкатегории
	DatetimeCreate datetime.Time `json:"datetimeCreate" validate:"required"` // Дата создания подкатегории
}

func (s *CreateTagReq) ConvertToRepoReq() repoModel.CreateTagReq {
	return repoModel.CreateTagReq{
		Name:            s.Name,
		AccountGroupID:  s.AccountGroupID,
		CreatedByUserID: s.Necessary.UserID,
		DatetimeCreate:  s.DatetimeCreate.Time,
	}
}

type UpdateTagReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32  `json:"id" validate:"required" minimum:"1"` // Идентификатор подкатегории
	Name      *string `json:"name" minimum:"1"`                   // Название подкатегории
}

type GetTagsReq struct {
	Necessary       services.NecessaryUserInformation
	AccountGroupIDs []uint32 `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"` // Идентификаторы групп счетов
}

type GetTagsToTransactionsReq struct {
	Necessary       services.NecessaryUserInformation
	AccountGroupIDs []uint32 `json:"-" schema:"-" minimum:"1"` // Идентификаторы групп счетов
	TransactionIDs  []uint32 `json:"-" schema:"-" minimum:"1"` // Идентификаторы транзакций
}

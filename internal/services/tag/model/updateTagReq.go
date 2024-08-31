package model

import "server/internal/services"

type UpdateTagReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32  `json:"id" validate:"required" minimum:"1"` // Идентификатор подкатегории
	Name      *string `json:"name" minimum:"1"`                   // Название подкатегории
}

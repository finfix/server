package model

import "server/internal/services"

type DeleteTagReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор подкатегории
}

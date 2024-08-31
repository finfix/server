package model

import (
	"server/internal/services"
)

type DeleteAccountGroupReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" schema:"id" validate:"required" minimum:"1"` // Идентификатор счета
}

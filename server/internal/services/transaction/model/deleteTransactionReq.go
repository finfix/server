package model

import (
	"pkg/necessary"
)

type DeleteTransactionReq struct {
	Necessary necessary.NecessaryUserInformation
	ID        uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

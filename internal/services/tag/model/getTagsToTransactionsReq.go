package model

import (
	"server/internal/services"
)

type GetTagsToTransactionsReq struct {
	Necessary       services.NecessaryUserInformation
	AccountGroupIDs []uint32 `json:"-" schema:"-" minimum:"1"` // Идентификаторы групп счетов
	TransactionIDs  []uint32 `json:"-" schema:"-" minimum:"1"` // Идентификаторы транзакций
}

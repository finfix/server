package model

import (
	"pkg/necessary"
)

type GetTagsToTransactionsReq struct {
	Necessary       necessary.NecessaryUserInformation
	AccountGroupIDs []uint32 `json:"-" schema:"-" minimum:"1"` // Идентификаторы групп счетов
	TransactionIDs  []uint32 `json:"-" schema:"-" minimum:"1"` // Идентификаторы транзакций
}

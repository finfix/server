package model

import (
	"server/app/pkg/datetime/date"
	"server/app/services/transaction/model/transactionType"
)

type CreateTransactionReq struct {
	Type            transactionType.Type
	AmountFrom      float64
	AmountTo        float64
	Note            string
	AccountFromID   uint32
	AccountToID     uint32
	DateTransaction date.Date
	IsExecuted      *bool
	CreatedByUserID uint32
}

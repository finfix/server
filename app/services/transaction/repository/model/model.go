package model

import (
	"github.com/shopspring/decimal"
	"time"

	"server/app/pkg/datetime"
	"server/app/services/transaction/model/transactionType"
)

type CreateTransactionReq struct {
	Type            transactionType.Type
	AmountFrom      decimal.Decimal
	AmountTo        decimal.Decimal
	Note            string
	AccountFromID   uint32
	AccountToID     uint32
	DateTransaction datetime.Date
	IsExecuted      *bool
	CreatedByUserID uint32
	DatetimeCreate  time.Time
}

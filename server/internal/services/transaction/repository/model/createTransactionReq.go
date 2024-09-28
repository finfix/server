package model

import (
	"time"

	"github.com/shopspring/decimal"

	"pkg/datetime"

	"server/internal/services/transaction/model/transactionType"
)

type CreateTransactionReq struct {
	Type               transactionType.Type
	AmountFrom         decimal.Decimal
	AmountTo           decimal.Decimal
	Note               string
	AccountFromID      uint32
	AccountToID        uint32
	DateTransaction    datetime.Date
	IsExecuted         bool
	CreatedByUserID    uint32
	DatetimeCreate     time.Time
	AccountingInCharts bool
}

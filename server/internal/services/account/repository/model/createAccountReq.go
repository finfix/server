package model

import (
	"time"

	"github.com/shopspring/decimal"

	"server/internal/services/account/model/accountType"
)

type CreateAccountReq struct {
	Budget             CreateReqBudget
	Name               string
	Visible            bool
	IconID             uint32
	Type               accountType.Type
	Currency           string
	AccountGroupID     uint32
	AccountingInHeader bool
	AccountingInCharts bool
	IsParent           bool
	ParentAccountID    *uint32
	UserID             uint32
	DatetimeCreate     time.Time
}

type CreateReqBudget struct {
	Amount         decimal.Decimal
	GradualFilling bool
	FixedSum       decimal.Decimal
	DaysOffset     uint32
}

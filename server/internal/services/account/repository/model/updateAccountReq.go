package model

import (
	"github.com/shopspring/decimal"
)

type UpdateAccountReq struct {
	Remainder          *decimal.Decimal
	Name               *string
	IconID             *uint32
	Visible            *bool
	AccountingInHeader *bool
	AccountingInCharts *bool
	Currency           *string
	ParentAccountID    *uint32
	SerialNumber       *uint32
	Budget             UpdateAccountBudgetReq
}

type UpdateAccountBudgetReq struct {
	Amount         *decimal.Decimal
	FixedSum       *decimal.Decimal
	DaysOffset     *uint32
	GradualFilling *bool
}

package model

import (
	"time"

	"server/app/pkg/datetime"
	"server/app/services/account/model/accountType"
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
	Amount         float64
	GradualFilling bool
	FixedSum       float64
	DaysOffset     uint32
}

type GetAccountsReq struct {
	IDs                []uint32
	AccountGroupIDs    []uint32
	Types              []accountType.Type
	AccountingInHeader *bool
	AccountingInCharts *bool
	Visible            *bool
	Currencies         []string
	IsParent           *bool
	ParentAccountIDs   []uint32
}

type CalculateRemaindersAccountsReq struct {
	IDs             []uint32
	AccountGroupIDs []uint32
	Types           []accountType.Type
	DateFrom        *datetime.Date
	DateTo          *datetime.Date
}

type UpdateAccountReq struct {
	Remainder          *float64
	Name               *string
	IconID             *uint32
	Visible            *bool
	AccountingInHeader *bool
	AccountingInCharts *bool
	Currency           *string
	ParentAccountID    *uint32
	Budget             UpdateAccountBudgetReq
}

type UpdateAccountBudgetReq struct {
	Amount         *float64
	FixedSum       *float64
	DaysOffset     *uint32
	GradualFilling *bool
}

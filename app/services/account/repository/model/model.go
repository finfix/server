package model

import (
	"server/app/pkg/datetime/date"
	"server/app/services/account/model/accountType"
)

type CreateReq struct {
	Budget          CreateReqBudget
	Name            string
	Visible         bool
	IconID          uint32
	Type            accountType.Type
	Currency        string
	AccountGroupID  uint32
	Accounting      bool
	IsParent        bool
	ParentAccountID *uint32
}

type CreateReqBudget struct {
	Amount         float64
	GradualFilling bool
	FixedSum       float64
	DaysOffset     uint32
}

type GetReq struct {
	IDs              []uint32
	AccountGroupIDs  []uint32
	Types            []accountType.Type
	Accounting       *bool
	Visible          *bool
	Currencies       []string
	IsParent         *bool
	ParentAccountIDs []uint32
}

type CalculateRemaindersAccountsReq struct {
	IDs             []uint32
	AccountGroupIDs []uint32
	Types           []accountType.Type
	DateFrom        *date.Date
	DateTo          *date.Date
}

type UpdateReq struct {
	Remainder       *float64
	Name            *string
	IconID          *uint32
	Visible         *bool
	Accounting      *bool
	Currency        *string
	ParentAccountID *uint32
	Budget          UpdateBudgetReq
}

type UpdateBudgetReq struct {
	Amount         *float64
	FixedSum       *float64
	DaysOffset     *uint32
	GradualFilling *bool
}

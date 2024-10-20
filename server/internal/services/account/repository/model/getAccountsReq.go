package model

import "server/internal/services/account/model/accountType"

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

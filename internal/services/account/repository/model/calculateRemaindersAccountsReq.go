package model

import (
	"server/internal/services/account/model/accountType"
	"server/pkg/datetime"
)

type CalculateRemaindersAccountsReq struct {
	IDs             []uint32
	AccountGroupIDs []uint32
	Types           []accountType.Type
	DateFrom        *datetime.Date
	DateTo          *datetime.Date
}

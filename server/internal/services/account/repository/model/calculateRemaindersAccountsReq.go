package model

import (
	"pkg/datetime"

	"server/internal/services/account/model/accountType"
)

type CalculateRemaindersAccountsReq struct {
	IDs             []uint32
	AccountGroupIDs []uint32
	Types           []accountType.Type
	DateFrom        *datetime.Date
	DateTo          *datetime.Date
}

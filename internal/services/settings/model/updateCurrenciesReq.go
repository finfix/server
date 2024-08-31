package model

import (
	"server/internal/services"
)

type UpdateCurrenciesReq struct {
	Necessary services.NecessaryUserInformation
}

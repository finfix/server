package model

import (
	"server/internal/services"
)

type SignOutReq struct {
	Necessary services.NecessaryUserInformation
}

package model

import (
	"server/app/pkg/datetime"
	"server/app/services"
)

type CreateReq struct {
	Name            string
	Email           string
	PasswordHash    string
	TimeCreate      datetime.Time
	DefaultCurrency string
}

type GetReq struct {
	Necessary services.NecessaryUserInformation
	IDs       []uint32
	Emails    []string
}

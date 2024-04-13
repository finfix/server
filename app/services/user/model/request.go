package model

import (
	"time"

	"server/app/services"
)

type CreateReq struct {
	Name            string
	Email           string
	PasswordHash    string
	TimeCreate      time.Time
	DefaultCurrency string
}

type GetReq struct {
	Necessary services.NecessaryUserInformation
	IDs       []uint32
	Emails    []string
}

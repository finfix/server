package model

import "time"

type GetReq struct {
	IDs    []uint32
	Emails []string
}

type CreateReq struct {
	Name            string
	Email           string
	PasswordHash    string
	TimeCreate      time.Time
	DefaultCurrency string
}

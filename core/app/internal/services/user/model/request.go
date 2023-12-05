package model

import "time"

type CreateReq struct {
	Name            string
	Email           string
	PasswordHash    string
	TimeCreate      time.Time
	DefaultCurrency string
}

type GetReq struct {
	IDs    []uint32
	Emails []string
}

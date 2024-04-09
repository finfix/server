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
	ID     uint32
	IDs    []uint32
	Emails []string
}

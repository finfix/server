package model

import (
	"time"

	"server/app/pkg/validation"
	"server/app/services"
)

type CreateReq struct {
	Name            string
	Email           string
	PasswordHash    []byte
	PasswordSalt    []byte
	TimeCreate      time.Time
	DefaultCurrency string
}

func (s CreateReq) Validate() error { return nil }

func (s CreateReq) SetNecessary(information services.NecessaryUserInformation) any {
	return s
}

type GetReq struct {
	Necessary services.NecessaryUserInformation
	IDs       []uint32
	Emails    []string
}

func (s GetReq) Validate() error {
	for _, email := range s.Emails {
		err := validation.Mail(email)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GetReq) SetNecessary(information services.NecessaryUserInformation) any {
	s.Necessary = information
	return s
}

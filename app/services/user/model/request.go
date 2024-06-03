package model

import (
	"time"

	"server/app/pkg/validation"
	"server/app/services"
	userRepoModel "server/app/services/user/repository/model"
)

type CreateReq struct {
	Name            string
	Email           string
	PasswordHash    []byte
	PasswordSalt    []byte
	TimeCreate      time.Time
	DefaultCurrency string
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

type UpdateUserReq struct {
	Necessary         services.NecessaryUserInformation
	Name              *string `json:"name"`
	Email             *string `json:"email"`
	Password          *string `json:"password"`
	OldPassword       *string `json:"oldPassword"`
	DefaultCurrency   *string `json:"defaultCurrency"`
	NotificationToken *string `json:"notificationToken"`
}

func (s UpdateUserReq) Validate() error {
	if s.Email != nil {
		err := validation.Mail(*s.Email)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s UpdateUserReq) SetNecessary(information services.NecessaryUserInformation) any {
	s.Necessary = information
	return s
}

func (s UpdateUserReq) ConvertToRepoModel() userRepoModel.UpdateUserReq {
	return userRepoModel.UpdateUserReq{
		ID:              s.Necessary.UserID,
		Name:            s.Name,
		Email:           s.Email,
		PasswordHash:    nil,
		PasswordSalt:    nil,
		DefaultCurrency: s.DefaultCurrency,
	}
}

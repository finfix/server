package model

import (
	"server/app/services"
	userModel "server/app/services/user/model"
)

type SendNotificationReq struct {
	Necessary    services.NecessaryUserInformation
	UserID       uint32                 `json:"userID"`
	Notification userModel.Notification `json:"notification"`
}

func (s SendNotificationReq) Validate() error {
	return nil
}

func (s SendNotificationReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	s.Necessary = necessary
	return s
}

type UpdateCurrenciesReq struct {
	Necessary services.NecessaryUserInformation
}

func (s UpdateCurrenciesReq) Validate() error {
	return nil
}

func (s UpdateCurrenciesReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	s.Necessary = necessary
	return s
}

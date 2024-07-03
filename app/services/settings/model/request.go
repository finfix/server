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

type UpdateCurrenciesReq struct {
	Necessary services.NecessaryUserInformation
}

func (s UpdateCurrenciesReq) Validate() error {
	return nil
}

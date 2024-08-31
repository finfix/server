package model

import (
	"server/internal/services"
	userModel "server/internal/services/user/model"
)

type SendNotificationReq struct {
	Necessary    services.NecessaryUserInformation
	UserID       uint32                 `json:"userID"`
	Notification userModel.Notification `json:"notification"`
}

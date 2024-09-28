package model

import (
	"pkg/necessary"
	userModel "server/internal/services/user/model"
)

type SendNotificationReq struct {
	Necessary    necessary.NecessaryUserInformation
	UserID       uint32                 `json:"userID"`
	Notification userModel.Notification `json:"notification"`
}

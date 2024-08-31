package service

import (
	"context"

	settingsModel "server/internal/services/settings/model"
)

func (s *Service) SendNotification(ctx context.Context, req settingsModel.SendNotificationReq) (res settingsModel.SendNotificationRes, err error) {

	err = s.checkAdmin(ctx, req.Necessary.UserID)
	if err != nil {
		return res, err
	}

	count, err := s.userService.SendNotification(
		ctx,
		req.UserID,
		req.Notification,
	)
	if err != nil {
		return res, err
	}
	res.NotificationsSent = count
	return res, err
}

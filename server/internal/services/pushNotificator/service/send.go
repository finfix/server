package service

import (
	"context"
	"time"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"

	"pkg/errors"
	"pkg/log"

	"server/internal/services/pushNotificator/model"
)

// SendNotification отправляет одно сообщение на все переданные устройства
func (s *PushNotificatorService) SendNotification(ctx context.Context, req model.SendNotificationReq) (id string, err error) {

	const defaultPriority = 5

	if !s.isOn {
		log.Warning(ctx, "Вызвана функция SendNotification. Пуши выключены")
		return id, nil
	}

	payload := payload.NewPayload()
	if req.Notification.Title != nil {
		payload = payload.AlertTitle(*req.Notification.Title)
	}
	if req.Notification.Subtitle != nil {
		payload = payload.AlertSubtitle(*req.Notification.Subtitle)
	}
	if req.Notification.Message != nil {
		payload = payload.AlertBody(*req.Notification.Message)
	}
	if req.Notification.Badge != nil {
		payload = payload.Badge(int(*req.Notification.Badge))
	}
	payload = payload.ContentAvailable()

	notification := &apns2.Notification{
		ApnsID:      "",
		CollapseID:  "",
		DeviceToken: req.NotificationToken,
		Topic:       req.BundleID,
		Expiration:  time.Time{},
		Priority:    defaultPriority,
		Payload:     payload,
		PushType:    apns2.PushTypeAlert,
	}

	res, err := s.apns.PushWithContext(ctx, notification)
	if err != nil {
		return id, errors.InternalServer.Wrap(err)
	}
	id = res.ApnsID

	if !res.Sent() {
		return id, errors.InternalServer.New(res.Reason)
	}

	return id, nil
}

package pushNotificator

import (
	"context"
	"time"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"

	"server/app/pkg/errors"
	"server/app/pkg/log"
)

type PushNotificator struct {
	apns *apns2.Client
	isOn bool
}

type APNsCredentials struct {
	TeamID      string
	KeyID       string
	KeyFilePath string
}

func NewPushNotificator(isOn bool, apnsCredentials APNsCredentials) (*PushNotificator, error) {

	if !isOn {
		log.Warning(context.Background(), "Push notificator is off")
		return &PushNotificator{
			isOn: isOn,
			apns: nil,
		}, nil
	}

	authKey, err := token.AuthKeyFromFile(apnsCredentials.KeyFilePath)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	apnsClient := apns2.NewTokenClient(&token.Token{ // nolint:exhaustruct
		AuthKey: authKey,
		KeyID:   apnsCredentials.KeyID,
		TeamID:  apnsCredentials.TeamID,
	})
	apnsClient.Host = apns2.HostProduction

	return &PushNotificator{
		isOn: isOn,
		apns: apnsClient,
	}, nil
}

// Push отправляет одно сообщение на все переданные устройства
func (s *PushNotificator) Push(ctx context.Context, req PushReq) (id string, err error) {

	const defaultPriority = 5

	if !s.isOn {
		log.Warning(ctx, "Вызвана функция Push. Пуши выключены")
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

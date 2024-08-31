package service

import (
	"context"

	"server/internal/services/user/model"
	"server/internal/services/user/model/OS"
	userRepoModel "server/internal/services/user/repository/model"
	"server/pkg/log"
	"server/pkg/pushNotificator"
)

// SendNotification отправляет пуш на все устройства пользователя
func (s *UserService) SendNotification(ctx context.Context, userID uint32, push model.Notification) (count uint8, err error) {

	// Получаем все девайсы пользователя
	devices, err := s.userRepository.GetDevices(ctx, userRepoModel.GetDevicesReq{ //nolint:exhaustruct
		UserIDs: []uint32{userID},
	})
	if err != nil {
		return count, err
	}

	// Проходимся по каждому девайсу
	for _, device := range devices {

		if device.NotificationToken == nil {
			continue
		}

		// Смотрим на операционную систему и отправляем уведомление
		switch device.DeviceInformation.NameOS {
		case OS.IOS, OS.IPadOS, OS.OSX, OS.WatchOS:
			_, err = s.pushNotificator.Push(ctx, pushNotificator.PushReq{
				Notification: pushNotificator.NotificationSettings{
					Title:    &push.Title,
					Message:  &push.Message,
					Subtitle: &push.Subtitle,
					Badge:    &push.BadgeCount,
				},
				NotificationToken: *device.NotificationToken,
				BundleID:          device.ApplicationInformation.BundleID,
			})

		case OS.Android:
			break
		}
		if err != nil {
			log.Error(ctx, err)
		}
		count++
	}

	return count, nil
}

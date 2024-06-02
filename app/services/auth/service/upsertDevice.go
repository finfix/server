package service

import (
	"context"

	userModel "server/app/services/user/model"
	userRepoModel "server/app/services/user/repository/model"
)

func (s *Service) upsertDevice(ctx context.Context, device userModel.Device) error {

	// Получаем девайс пользователя
	devices, err := s.userRepository.GetDevices(ctx, userRepoModel.GetDevicesReq{ //nolint:exhaustruct
		DeviceIDs: []string{device.DeviceID},
		UserIDs:   []uint32{device.UserID},
	})
	if err != nil {
		return err
	}

	// Если девайс не нашелся
	if len(devices) == 0 {

		_, err = s.userRepository.CreateDevice(ctx, device)

	} else { // Если девайс нашелся

		// Обновляем у него токен
		err = s.userRepository.UpdateDevice(ctx, userRepoModel.UpdateDeviceReq{
			DeviceID:          device.DeviceID,
			UserID:            device.UserID,
			RefreshToken:      &device.RefreshToken,
			NotificationToken: nil,
			ApplicationInformation: userRepoModel.UpdateApplicationInformationReq{
				BundleID: &device.ApplicationInformation.BundleID,
				Version:  &device.ApplicationInformation.Version,
				Build:    &device.ApplicationInformation.Build,
			},
			DeviceInformation: userRepoModel.UpdateDeviceInformationReq{
				VersionOS: &device.DeviceInformation.VersionOS,
			},
		})
	}

	if err != nil {
		return err
	}

	return nil
}

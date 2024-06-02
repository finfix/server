package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/pkg/jwtManager"
	"server/app/services/auth/model"
	userRepoModel "server/app/services/user/repository/model"
)

// RefreshTokens обновляет токены доступа в базе данных
func (s *Service) RefreshTokens(ctx context.Context, req model.RefreshTokensReq) (newTokens model.RefreshTokensRes, err error) {

	// Получаем девайс по идентификатору пользователя и девайса
	devices, err := s.userRepository.GetDevices(ctx, userRepoModel.GetDevicesReq{ // nolint:exhaustruct
		DeviceIDs: []string{req.Necessary.DeviceID},
		UserIDs:   []uint32{req.Necessary.UserID},
	})
	if err != nil {
		return newTokens, err
	}
	if len(devices) == 0 {
		return newTokens, errors.Unauthorized.New("Device not found", []errors.Option{
			errors.HumanTextOption("Девайс не найдена"),
		}...)
	}
	device := devices[0]

	// Сравниваем токен из базы данных с переданным пользователем токеном
	if req.Token != device.RefreshToken {
		return newTokens, errors.Forbidden.New("Token is incorrect")
	}

	// Смотрим, не истек ли еще токен
	userID, deviceID, err := jwtManager.Parse(device.RefreshToken)
	if err != nil {
		return newTokens, err
	}

	// Дополнительно проверяем идентификаторы
	if userID != req.Necessary.UserID {
		return newTokens, errors.Forbidden.New("UserID not matched")
	}
	if deviceID != req.Necessary.DeviceID {
		return newTokens, errors.Forbidden.New("DeviceID not matched")
	}

	// Создаем новую пару токенов
	newTokens.Tokens, err = s.createPairTokens(ctx, req.Necessary.UserID, req.Necessary.DeviceID)
	if err != nil {
		return newTokens, err
	}

	// Обновляем refresh токен у девайса
	if err = s.userRepository.UpdateDevice(ctx, userRepoModel.UpdateDeviceReq{
		UserID:            req.Necessary.UserID,
		DeviceID:          req.Necessary.DeviceID,
		RefreshToken:      &newTokens.Tokens.RefreshToken,
		NotificationToken: nil,
	}); err != nil {
		return newTokens, err
	}

	// Возвращаем пару токенов клиенту
	return newTokens, nil
}

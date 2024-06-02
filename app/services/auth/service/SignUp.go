package service

import (
	"context"
	"time"

	"server/app/pkg/errors"
	"server/app/pkg/hasher"
	"server/app/services/auth/model"
	userModel "server/app/services/user/model"
)

// SignUp регистрирует пользователя и возвращает токены доступа
func (s *Service) SignUp(ctx context.Context, loginData model.SignUpReq) (accessData model.AuthRes, err error) {

	// Проверяем, есть ли пользователь в бд с таким email
	if _users, err := s.userRepository.GetUsers(ctx, userModel.GetReq{Emails: []string{loginData.Email}}); err != nil { //nolint:exhaustruct
		return accessData, err
	} else if len(_users) != 0 {
		return accessData, errors.Forbidden.New("User with this email is already registered", []errors.Option{
			errors.HumanTextOption("Пользователь с таким email уже зарегистрирован"),
			errors.ParamsOption("email", loginData.Email),
		}...)
	}

	// Получаем хэш пароля пользователя
	passwordHash, passwordSalt, err := hasher.CreateNewPassword([]byte(loginData.Password), s.generalSalt)
	if err != nil {
		return accessData, err
	}

	return accessData, s.generalRepository.WithinTransaction(ctx, func(ctx context.Context) error {

		// Создаем пользователя
		accessData.ID, err = s.userRepository.CreateUser(ctx, userModel.CreateReq{
			Name:            loginData.Name,
			Email:           loginData.Email,
			PasswordHash:    passwordHash,
			PasswordSalt:    passwordSalt,
			TimeCreate:      time.Now(),
			DefaultCurrency: "RUB", // TODO: Поменять
		})
		if err != nil {
			return err
		}

		// Создаем пару токенов
		accessData.Tokens, err = s.createPairTokens(ctx, accessData.ID, loginData.DeviceID)
		if err != nil {
			return err
		}

		// Создаем или обновляем девайс пользователя
		err = s.upsertDevice(ctx, userModel.Device{
			DeviceInformation:      loginData.Device,
			NotificationToken:      nil,
			RefreshToken:           accessData.RefreshToken,
			UserID:                 accessData.ID,
			DeviceID:               loginData.DeviceID,
			ApplicationInformation: loginData.Application,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/pkg/hasher"
	"server/app/services/auth/model"
	userModel "server/app/services/user/model"
)

// SignIn авторизует пользователя и возвращает токены доступа
func (s *Service) SignIn(ctx context.Context, loginData model.SignInReq) (accessData model.AuthRes, err error) {

	// Получаем пользователя по email
	users, err := s.userRepository.GetUsers(ctx, userModel.GetReq{Emails: []string{loginData.Email}}) //nolint:exhaustruct
	if err != nil {
		return accessData, err
	}
	if len(users) == 0 {
		return accessData, errors.NotFound.New("User not found", []errors.Option{
			errors.HumanTextOption("Пользователь не найден"),
		}...)
	}
	user := users[0]

	accessData.ID = user.ID

	// Сравниваем пришедший пароль и хэш пароля из базы данных
	if err = hasher.CompareHashAndPassword(user.PasswordHash, []byte(loginData.Password), user.PasswordSalt, s.generalSalt); err != nil {
		return accessData, err
	}

	// Создаем пару токенов
	accessData.Tokens, err = s.createPairTokens(ctx, user.ID, loginData.DeviceID)
	if err != nil {
		return accessData, err
	}

	// Создаем или обновляем девайс пользователя
	err = s.upsertDevice(ctx, userModel.Device{
		OS:                loginData.OS,
		NotificationToken: nil,
		RefreshToken:      accessData.RefreshToken,
		UserID:            accessData.ID,
		DeviceID:          loginData.DeviceID,
		BundleID:          loginData.BundleID,
	})
	if err != nil {
		return accessData, err
	}

	// Возвращаем идентификатор пользователя и токены
	return accessData, nil
}

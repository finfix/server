package service

import (
	"context"

	"pkg/errors"
	"pkg/passwordManager"

	"server/internal/services/auth/model"
	"server/internal/services/auth/service/utils"
	userModel "server/internal/services/user/model"
)

// SignIn авторизует пользователя и возвращает токены доступа
func (s *AuthService) SignIn(ctx context.Context, loginData model.SignInReq) (accessData model.AuthRes, err error) {

	// Получаем пользователя по email
	users, err := s.userRepository.GetUsers(ctx, userModel.GetUsersReq{Emails: []string{loginData.Email}}) //nolint:exhaustruct
	if err != nil {
		return accessData, err
	}
	if len(users) == 0 {
		return accessData, errors.NotFound.New("User not found",
			errors.HumanTextOption("Пользователь не найден"),
		)
	}
	user := users[0]

	accessData.ID = user.ID

	// Сравниваем пришедший пароль и хэш пароля из базы данных
	if err = passwordManager.CompareHashAndPassword(user.PasswordHash, []byte(loginData.Password), user.PasswordSalt, s.generalSalt); err != nil {
		return accessData, err
	}

	// Создаем пару токенов
	accessData.Tokens, err = utils.CreatePairTokens(ctx, user.ID, loginData.DeviceID)
	if err != nil {
		return accessData, err
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
		return accessData, err
	}

	// Возвращаем идентификатор пользователя и токены
	return accessData, nil
}

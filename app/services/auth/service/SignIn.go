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
	users, err := s.userService.GetUsers(ctx, userModel.GetReq{Emails: []string{loginData.Email}})
	if err != nil {
		return accessData, err
	}
	if len(users) == 0 {
		return accessData, errors.NotFound.New("User not found", errors.Options{
			HumanText: "Пользователь не найден",
		})
	}
	user := users[0]

	// Сравниваем пришедший пароль и хэш пароля из базы данных
	if err = hasher.CompareHashAndPassword(user.PasswordHash, []byte(loginData.Password), user.PasswordSalt, s.generalSalt); err != nil {
		return accessData, err
	}

	// Создаем сессию
	accessData.AccessToken, accessData.RefreshToken, err = s.createSession(ctx, user.ID, loginData.DeviceID)
	if err != nil {
		return accessData, err
	}

	accessData.ID = user.ID

	return accessData, nil
}

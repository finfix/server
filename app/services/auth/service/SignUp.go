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
// TODO: Добавить SQL-транзакцию
func (s *Service) SignUp(ctx context.Context, user model.SignUpReq) (accessData model.AuthRes, err error) {

	// Проверяем, есть ли пользователь в бд с таким email
	if _users, err := s.userService.GetUsers(ctx, userModel.GetReq{Emails: []string{user.Email}}); err != nil { //nolint:exhaustruct
		return accessData, err
	} else if len(_users) != 0 {
		return accessData, errors.Forbidden.New("User with this email is already registered", []errors.Option{
			errors.HumanTextOption("Пользователь с таким email уже зарегистрирован"),
			errors.ParamsOption("email", user.Email),
		}...)
	}

	// Получаем хэш пароля пользователя
	passwordHash, passwordSalt, err := hasher.CreateNewPassword([]byte(user.Password), s.generalSalt)
	if err != nil {
		return accessData, err
	}

	return accessData, s.generalRepository.WithinTransaction(ctx, func(ctx context.Context) error {

		// Заносим пользователя в базу данных
		accessData.ID, err = s.userService.CreateUser(ctx, userModel.CreateReq{
			Name:            user.Name,
			Email:           user.Email,
			PasswordHash:    passwordHash,
			PasswordSalt:    passwordSalt,
			TimeCreate:      time.Now(),
			DefaultCurrency: "RUB", // TODO: Поменять
		})
		if err != nil {
			return err
		}

		// Создаем сессию
		accessData.AccessToken, accessData.RefreshToken, err = s.createSession(ctx, accessData.ID, user.DeviceID)
		if err != nil {
			return err
		}

		return nil
	})
}

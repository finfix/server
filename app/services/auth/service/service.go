package service

import (
	"context"
	"time"

	"server/app/config"
	"server/app/pkg/auth"
	"server/app/pkg/datetime"
	"server/app/pkg/errors"
	"server/app/pkg/hasher"
	"server/app/pkg/logging"
	model3 "server/app/services/auth/model"
	authRepository "server/app/services/auth/repository"
	model2 "server/app/services/user/model"
	userService "server/app/services/user/service"
)

var _ AuthRepository = &authRepository.Repository{}
var _ UserService = &userService.Service{}

type AuthRepository interface {
	CreateSession(ctx context.Context, id uint32, deviceID string) (string, error)
	DeleteSession(ctx context.Context, oldRefreshToken string) error
	GetSession(context.Context, string) (uint32, string, error)
}

type UserService interface {
	GetTransactions(context.Context, model2.GetReq) ([]model2.User, error)
	CreateUser(context.Context, model2.CreateReq) (uint32, error)
}

// SignIn авторизует пользователя и возвращает токены доступа
func (s *Service) SignIn(ctx context.Context, loginData model3.SignInReq) (accessData model3.AuthRes, err error) {

	// Получаем идентификатор пользователя
	users, err := s.user.GetTransactions(ctx, model2.GetReq{Emails: []string{loginData.Email}})
	if err != nil {
		return accessData, err
	}

	if len(users) == 0 {
		return accessData, errors.NotFound.New("User not found", errors.Options{
			HumanText: "Пользователь не найден",
		})
	}
	user := users[0]

	// Получаем хэш пароля пользователя
	passwordHash, err := hasher.Hash(loginData.Password, config.GetConfig().SHASalt)
	if err != nil {
		return accessData, err
	}

	// Сравниваем пароль пользователя с паролем из бд
	if user.PasswordHash != passwordHash {
		return accessData, errors.BadRequest.New("Incorrect password or login", errors.Options{
			HumanText: "Неверно введен логин или пароль",
		})
	}

	// Создаем сессию
	accessData.AccessToken, accessData.RefreshToken, err = s.createSession(ctx, user.ID, loginData.DeviceID)
	if err != nil {
		return accessData, err
	}

	accessData.ID = user.ID

	return accessData, nil
}

// SignUp регистрирует пользователя и возвращает токены доступа
// TODO: Добавить SQL-транзакцию
func (s *Service) SignUp(ctx context.Context, user model3.SignUpReq) (accessData model3.AuthRes, err error) {

	// Проверяем, есть ли пользователь в бд
	if _users, err := s.user.GetTransactions(ctx, model2.GetReq{Emails: []string{user.Email}}); err != nil {
		return accessData, err
	} else if len(_users) != 0 {
		return accessData, errors.Forbidden.New("User with this email is already registered", errors.Options{
			HumanText: "Пользователь с таким email уже зарегистрирован",
			Params: map[string]any{
				"email": user.Email,
			},
		})
	}

	// Получаем хэш пароля пользователя
	user.Password, err = hasher.Hash(user.Password, config.GetConfig().SHASalt)
	if err != nil {
		return accessData, err
	}

	// Заносим пользователя в базу данных
	accessData.ID, err = s.user.CreateUser(ctx, model2.CreateReq{
		Name:            user.Name,
		Email:           user.Email,
		PasswordHash:    user.Password,
		TimeCreate:      datetime.Time{Time: time.Now()},
		DefaultCurrency: "RUB", // TODO: Поменять
	})
	if err != nil {
		return accessData, err
	}

	// Создаем сессию
	accessData.AccessToken, accessData.RefreshToken, err = s.createSession(ctx, accessData.ID, user.DeviceID)
	if err != nil {
		return accessData, err
	}

	return accessData, nil
}

// createSession создает токены доступа и заносит сессию в базу данных
func (s *Service) createSession(ctx context.Context, id uint32, deviceID string) (accessToken, refreshToken string, err error) {

	// Получаем время жизни токена
	durationAccess, err := time.ParseDuration(config.GetConfig().Token.AccessTokenTTL)
	if err != nil {
		return "", "", err
	}

	signingKey := config.GetConfig().Token.SigningKey

	// Создаем Access token
	accessToken, err = auth.NewJWT(id, signingKey, deviceID, durationAccess)
	if err != nil {
		return "", "", err
	}

	// Создаем и заносим новую сессию в базу данных
	refreshToken, err = s.auth.CreateSession(ctx, id, deviceID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// RefreshTokens обновляет токены доступа в базе данных
// TODO: Добавить SQL-транзакцию
func (s *Service) RefreshTokens(ctx context.Context, refreshToken string) (newTokens model3.RefreshTokensRes, err error) {

	// Смотрим, есть ли сессия и валидируем ее
	id, deviceID, err := s.auth.GetSession(ctx, refreshToken)
	if err != nil {
		return newTokens, err
	}

	// Удаляем старую сессию в бд
	if err = s.auth.DeleteSession(ctx, refreshToken); err != nil {
		return newTokens, err
	}

	// Создаем новую сессию и получаем новые токены
	newTokens.AccessToken, newTokens.RefreshToken, err = s.createSession(ctx, id, deviceID)
	if err != nil {
		return newTokens, err
	}

	return newTokens, nil
}

type Service struct {
	auth   AuthRepository
	user   UserService
	logger *logging.Logger
}

func New(
	auth AuthRepository,
	user UserService,
	logger *logging.Logger,
) *Service {
	return &Service{
		auth:   auth,
		user:   user,
		logger: logger,
	}
}

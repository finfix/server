package service

import (
	"context"
	"time"

	"server/app/pkg/auth"
	"server/app/pkg/logging"
	authModel "server/app/services/auth/model"
	authRepository "server/app/services/auth/repository"
	"server/app/services/generalRepository"
	userModel "server/app/services/user/model"
	userService "server/app/services/user/service"
)

var _ AuthRepository = &authRepository.Repository{}
var _ UserService = &userService.Service{}
var _ GeneralRepository = &generalRepository.Repository{}

type AuthRepository interface {
	CreateSession(ctx context.Context, token string, timeExpiry time.Time, deviceID string, userID uint32) error
	DeleteSession(ctx context.Context, userID uint32, deviceID string) error
	GetSession(context.Context, authModel.RefreshTokensReq) (authModel.Session, error)
}

type UserService interface {
	GetUsers(context.Context, userModel.GetReq) ([]userModel.User, error)
	CreateUser(context.Context, userModel.CreateReq) (uint32, error)
}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}

func (s *Service) createSession(ctx context.Context, userID uint32, deviceID string) (accessToken, refreshToken string, err error) {

	// Создаем Access token
	accessToken, err = auth.NewJWT(userID, deviceID)
	if err != nil {
		return "", "", err
	}

	// Создаем refresh token
	refreshToken, refreshTokenExpiresAt, err := auth.NewRefreshToken()
	if err != nil {
		return "", "", err
	}

	// Создаем и заносим новую сессию в базу данных
	err = s.authRepository.CreateSession(ctx, refreshToken, refreshTokenExpiresAt, deviceID, userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

type Service struct {
	authRepository    AuthRepository
	userService       UserService
	generalRepository GeneralRepository
	generalSalt       []byte
	logger            *logging.Logger
}

func New(
	authRepository AuthRepository,
	userService UserService,
	generalRepository GeneralRepository,
	generalSalt []byte,
	logger *logging.Logger,
) *Service {
	return &Service{
		authRepository:    authRepository,
		userService:       userService,
		generalRepository: generalRepository,
		generalSalt:       generalSalt,
		logger:            logger,
	}
}

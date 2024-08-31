package service

import (
	"context"

	"server/internal/services/generalRepository"
	userModel "server/internal/services/user/model"
	userRepository "server/internal/services/user/repository"
	userRepoModel "server/internal/services/user/repository/model"
)

var _ UserRepository = &userRepository.UserRepository{}
var _ GeneralRepository = &generalRepository.GeneralRepository{}

type UserRepository interface {
	GetUsers(context.Context, userModel.GetUsersReq) ([]userModel.User, error)
	CreateUser(context.Context, userModel.CreateReq) (uint32, error)

	CreateDevice(context.Context, userModel.Device) (uint32, error)
	DeleteDevice(ctx context.Context, userID uint32, deviceID string) error
	UpdateDevice(context.Context, userRepoModel.UpdateDeviceReq) error
	GetDevices(context.Context, userRepoModel.GetDevicesReq) ([]userModel.Device, error)
}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}

type AuthService struct {
	userRepository    UserRepository
	generalRepository GeneralRepository
	generalSalt       []byte
}

func NewAuthService(
	userRepository UserRepository,
	generalRepository GeneralRepository,
	generalSalt []byte,

) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		generalRepository: generalRepository,
		generalSalt:       generalSalt,
	}
}

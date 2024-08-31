package service

import (
	"context"

	"server/app/services/generalRepository"
	userModel "server/app/services/user/model"
	userRepository "server/app/services/user/repository"
	userRepoModel "server/app/services/user/repository/model"
)

var _ UserRepository = &userRepository.Repository{}
var _ GeneralRepository = &generalRepository.Repository{}

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

type Service struct {
	userRepository    UserRepository
	generalRepository GeneralRepository
	generalSalt       []byte
}

func New(
	userRepository UserRepository,
	generalRepository GeneralRepository,
	generalSalt []byte,

) *Service {
	return &Service{
		userRepository:    userRepository,
		generalRepository: generalRepository,
		generalSalt:       generalSalt,
	}
}

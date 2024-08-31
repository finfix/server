package service

import (
	"context"

	"server/internal/services/generalRepository"
	userModel "server/internal/services/user/model"
	userRepository "server/internal/services/user/repository"
	userRepoModel "server/internal/services/user/repository/model"
	"server/pkg/pushNotificator"
)

var _ UserRepository = &userRepository.UserRepository{}
var _ GeneralRepository = &generalRepository.GeneralRepository{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
}

type UserRepository interface {
	CreateUser(context.Context, userModel.CreateReq) (uint32, error)
	GetUsers(context.Context, userModel.GetUsersReq) ([]userModel.User, error)
	UpdateUser(context.Context, userRepoModel.UpdateUserReq) error

	LinkUserToAccountGroup(context.Context, uint32, uint32) error

	GetDevices(context.Context, userRepoModel.GetDevicesReq) ([]userModel.Device, error)
	UpdateDevice(context.Context, userRepoModel.UpdateDeviceReq) error
}

type UserService struct {
	userRepository    UserRepository
	generalRepository GeneralRepository
	pushNotificator   *pushNotificator.PushNotificator
	generalSalt       []byte
}

func NewUserService(
	userRepository UserRepository,
	generalRepository GeneralRepository,
	pushNotificator *pushNotificator.PushNotificator,
	generalSalt []byte,
) *UserService {
	return &UserService{
		userRepository:    userRepository,
		generalRepository: generalRepository,
		pushNotificator:   pushNotificator,
		generalSalt:       generalSalt,
	}
}

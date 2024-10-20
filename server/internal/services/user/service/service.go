package service

import (
	"context"

	pushNotificatorModel "server/internal/services/pushNotificator/model"
	pushNotificatorService "server/internal/services/pushNotificator/service"
	"server/internal/services/transactor"
	userModel "server/internal/services/user/model"
	userRepository "server/internal/services/user/repository"
	userRepoModel "server/internal/services/user/repository/model"
)

var _ UserRepository = new(userRepository.UserRepository)
var _ GeneralRepository = new(transactor.Transactor)
var _ PushNotificatorService = new(pushNotificatorService.PushNotificatorService)

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

	GetAccessedAccountGroups(ctx context.Context, userID uint32) ([]uint32, error)
}

type PushNotificatorService interface {
	SendNotification(ctx context.Context, req pushNotificatorModel.SendNotificationReq) (string, error)
}

type UserService struct {
	userRepository    UserRepository
	generalRepository GeneralRepository
	pushNotificator   PushNotificatorService
	generalSalt       []byte
}

func NewUserService(
	userRepository UserRepository,
	generalRepository GeneralRepository,
	pushNotificator PushNotificatorService,
	generalSalt []byte,
) *UserService {
	return &UserService{
		userRepository:    userRepository,
		generalRepository: generalRepository,
		pushNotificator:   pushNotificator,
		generalSalt:       generalSalt,
	}
}

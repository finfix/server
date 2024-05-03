package service

import (
	"context"

	accountModel "server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository"
	userModel "server/app/services/user/model"
	userRepository "server/app/services/user/repository"
)

var _ UserRepository = &userRepository.Repository{}
var _ GeneralRepository = &generalRepository.Repository{}

type UserRepository interface {
	CreateUser(context.Context, userModel.CreateReq) (uint32, error)
	GetUsers(context.Context, userModel.GetReq) ([]userModel.User, error)
	LinkUserToAccountGroup(context.Context, uint32, uint32) error
}

type AccountRepository interface {
	CreateAccountGroup(context.Context, accountModel.CreateAccountGroupReq) (uint32, error)
	CreateAccount(ctx context.Context, req accountRepoModel.CreateAccountReq) (uint32, uint32, error)
}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
}

type Service struct {
	user    UserRepository
	account AccountRepository
	general GeneralRepository
}

// CreateUser создает нового пользователя
func (s *Service) CreateUser(ctx context.Context, user userModel.CreateReq) (id uint32, err error) {
	return s.user.CreateUser(ctx, user)
}

// GetUsers возвращает всех юзеров по фильтрам
func (s *Service) GetUsers(ctx context.Context, filters userModel.GetReq) (users []userModel.User, err error) {
	return s.user.GetUsers(ctx, filters)
}

func New(
	user UserRepository,
	general GeneralRepository,
	account AccountRepository,

) *Service {
	return &Service{
		user:    user,
		general: general,
		account: account,
	}
}

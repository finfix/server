package service

import (
	"context"

	userModel "server/app/services/user/model"
	userRepository "server/app/services/user/repository"
)

var _ UserRepository = &userRepository.Repository{}

type UserRepository interface {
	CreateUser(context.Context, userModel.CreateReq) (uint32, error)
	GetUsers(context.Context, userModel.GetReq) ([]userModel.User, error)
	LinkUserToAccountGroup(context.Context, uint32, uint32) error
}

type Service struct {
	userRepository UserRepository
}

// CreateUser создает нового пользователя
func (s *Service) CreateUser(ctx context.Context, user userModel.CreateReq) (id uint32, err error) {
	return s.userRepository.CreateUser(ctx, user)
}

// GetUsers возвращает всех юзеров по фильтрам
func (s *Service) GetUsers(ctx context.Context, filters userModel.GetReq) (users []userModel.User, err error) {
	return s.userRepository.GetUsers(ctx, filters)
}

func New(
	userRepository UserRepository,
) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

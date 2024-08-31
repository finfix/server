package service

import (
	"context"

	"server/internal/services/user/model"
)

// CreateUser создает нового пользователя
func (s *UserService) CreateUser(ctx context.Context, user model.CreateReq) (id uint32, err error) {
	return s.userRepository.CreateUser(ctx, user)
}

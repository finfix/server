package service

import (
	"context"

	"server/internal/services/user/model"
)

// GetUsers возвращает всех юзеров по фильтрам
func (s *Service) GetUsers(ctx context.Context, filters model.GetUsersReq) (users []model.User, err error) {
	return s.userRepository.GetUsers(ctx, filters)
}

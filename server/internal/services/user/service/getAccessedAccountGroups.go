package service

import (
	"context"
)

// GetAccessedAccountGroups возвращает доступы пользователей к группам счетов
func (s *UserService) GetAccessedAccountGroups(ctx context.Context, userID uint32) ([]uint32, error) {
	return s.userRepository.GetAccessedAccountGroups(ctx, userID)
}

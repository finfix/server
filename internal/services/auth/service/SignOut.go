package service

import (
	"context"

	"server/internal/services/auth/model"
)

// SignOut удаляет данные девайса пользователя
func (s *Service) SignOut(ctx context.Context, req model.SignOutReq) error {
	return s.userRepository.DeleteDevice(ctx, req.Necessary.UserID, req.Necessary.DeviceID)
}

package service

import (
	"context"

	userModel "server/internal/services/user/model"
	"server/pkg/errors"
)

func (s *SettingsService) checkAdmin(ctx context.Context, userID uint32) error {
	users, err := s.userService.GetUsers(ctx, userModel.GetUsersReq{ //nolint:exhaustruct
		IDs: []uint32{userID},
	})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.NotFound.New("User not found")
	}
	user := users[0]
	if !user.IsAdmin {
		return errors.Forbidden.New("Access denied")
	}
	return nil
}

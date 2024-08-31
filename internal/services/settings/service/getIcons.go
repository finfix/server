package service

import (
	"context"

	settingsModel "server/internal/services/settings/model"
)

func (s *Service) GetIcons(ctx context.Context) ([]settingsModel.Icon, error) {
	return s.settingsRepository.GetIcons(ctx)
}

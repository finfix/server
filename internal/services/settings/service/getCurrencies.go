package service

import (
	"context"

	settingsModel "server/internal/services/settings/model"
)

func (s *Service) GetCurrencies(ctx context.Context) ([]settingsModel.Currency, error) {
	return s.settingsRepository.GetCurrencies(ctx)
}

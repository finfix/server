package service

import (
	"context"

	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/model/applicationType"
	"server/pkg/errors"
)

func (s *Service) GetVersion(ctx context.Context, appType applicationType.Type) (version settingsModel.Version, err error) {
	switch appType {
	case applicationType.Server:
		return settingsModel.Version{
			Version: s.version.Version,
			Build:   s.version.Build,
		}, nil
	case applicationType.IOs:
		return s.settingsRepository.GetVersion(ctx, appType)
	case applicationType.Android, applicationType.Web:
		return version, errors.NotFound.New("Такое приложение еще не реализовано")
	default:
		return version, errors.BadRequest.New("Неверный тип приложения")
	}
}

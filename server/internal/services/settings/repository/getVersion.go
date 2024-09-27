package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/model/applicationType"
)

func (repo *SettingsRepository) GetVersion(ctx context.Context, appType applicationType.Type) (version settingsModel.Version, err error) {
	return version, repo.db.Get(ctx, &version, sq.
		Select("*").
		From("settings.versions").
		Where("name = ?", appType).
		Limit(1),
	)
}

package repository

import (
	"context"

	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/model/applicationType"
)

func (repo *SettingsRepository) GetVersion(ctx context.Context, appType applicationType.Type) (version settingsModel.Version, err error) {
	return version, repo.db.Get(ctx, &version, `
			SELECT * 
			FROM settings.versions 
			WHERE name = ?
			LIMIT 1`,
		appType,
	)
}

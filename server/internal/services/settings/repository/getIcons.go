package repository

import (
	"context"

	settingsModel "server/internal/services/settings/model"
)

func (repo *SettingsRepository) GetIcons(ctx context.Context) (icons []settingsModel.Icon, err error) {
	return icons, repo.db.Select(ctx, &icons, `SELECT * FROM coin.icons`)
}

package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	settingsModel "server/internal/services/settings/model"
)

func (repo *SettingsRepository) GetIcons(ctx context.Context) (icons []settingsModel.Icon, err error) {
	return icons, repo.db.Select(ctx, &icons, sq.
		Select("*").
		From("coin.icons"),
	)
}

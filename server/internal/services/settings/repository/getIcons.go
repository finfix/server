package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/repository/iconDDL"
)

func (repo *SettingsRepository) GetIcons(ctx context.Context) (icons []settingsModel.Icon, err error) {
	return icons, repo.db.Select(ctx, &icons, sq.
		Select(ddlHelper.SelectAll).
		From(iconDDL.Table),
	)
}

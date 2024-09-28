package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/repository/currencyDDL"
)

func (repo *SettingsRepository) GetCurrencies(ctx context.Context) (currencies []settingsModel.Currency, err error) {
	return currencies, repo.db.Select(ctx, &currencies, sq.
		Select(ddlHelper.SelectAll).
		From(currencyDDL.Table),
	)
}

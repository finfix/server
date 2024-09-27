package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	settingsModel "server/internal/services/settings/model"
)

func (repo *SettingsRepository) GetCurrencies(ctx context.Context) (currencies []settingsModel.Currency, err error) {
	return currencies, repo.db.Select(ctx, &currencies, sq.
		Select("*").
		From("coin.currencies"),
	)
}

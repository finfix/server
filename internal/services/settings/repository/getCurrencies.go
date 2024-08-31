package repository

import (
	"context"

	settingsModel "server/internal/services/settings/model"
)

func (repo *SettingsRepository) GetCurrencies(ctx context.Context) (currencies []settingsModel.Currency, err error) {
	return currencies, repo.db.Select(ctx, &currencies, `SELECT * FROM coin.currencies`)
}

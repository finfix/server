package generalRepository

import (
	"context"

	"github.com/shopspring/decimal"
)

// GetCurrencies возвращает список валют и их курсов
func (repo *GeneralRepository) GetCurrencies(ctx context.Context) (map[string]decimal.Decimal, error) {
	var currencies []struct {
		Name string          `db:"signatura"`
		Rate decimal.Decimal `db:"rate"`
	}
	if err := repo.db.Select(ctx, &currencies, `
			SELECT * 
			FROM coin.currencies`,
	); err != nil {
		return nil, err
	}

	rates := make(map[string]decimal.Decimal, len(currencies))
	for _, currency := range currencies {
		rates[currency.Name] = currency.Rate
	}

	return rates, nil
}

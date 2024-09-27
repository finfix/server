package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/shopspring/decimal"
)

// UpdateCurrencies обновляет курсы валют в базе данных
func (repo *SettingsRepository) UpdateCurrencies(ctx context.Context, rates map[string]decimal.Decimal) error {
	var (
		args = make([]any, 0, len(rates)*2) //nolint:gomnd // 2 - количество аргументов на одну запись (signatura и rate)
	)

	// Формируем аргументы для запроса
	for currency, rate := range rates {
		args = append(args, currency, currency, rate, currency)
	}

	q := sq.Insert("coin.currencies").
		Columns("signatura", "name", "rate", "symbol").
		Values(args...).
		Suffix("ON CONFLICT (signatura) DO UPDATE SET rate = EXCLUDED.rate")

	// Обновляем курсы валют
	return repo.db.Exec(ctx, q)
}

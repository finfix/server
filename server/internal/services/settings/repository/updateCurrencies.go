package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/shopspring/decimal"

	"server/internal/services/settings/repository/currencyDDL"
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

	q := sq.Insert(currencyDDL.Table).
		Columns(currencyDDL.ColumnSlug, currencyDDL.ColumnName, currencyDDL.ColumnRate, currencyDDL.ColumnSymbol).
		Values(args...).
		Suffix(fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.%s", currencyDDL.ColumnSlug, currencyDDL.ColumnRate, currencyDDL.ColumnRate))

	// Обновляем курсы валют
	return repo.db.Exec(ctx, q)
}

package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// UpdateCurrencies обновляет курсы валют в базе данных
func (repo *Repository) UpdateCurrencies(ctx context.Context, rates map[string]decimal.Decimal) error {
	var (
		pattern  = "(?, ?, ?, ?)"
		tmpQuery = make([]string, 0, len(rates))
		args     = make([]interface{}, 0, len(rates)*2) //nolint:gomnd // 2 - количество аргументов на одну запись (signatura и rate)
	)

	// Формируем аргументы для запроса
	for currency, rate := range rates {
		tmpQuery = append(tmpQuery, pattern)
		args = append(args, currency, currency, rate, currency)
	}

	// Конструируем запрос
	query := fmt.Sprintf(`
			INSERT INTO coin.currencies (signatura, name, rate, symbol)
			VALUES %v
			ON CONFLICT (signatura)
			DO UPDATE 
			  SET rate = EXCLUDED.rate`,
		strings.Join(tmpQuery, ", "))

	// Обновляем курсы валют
	return repo.db.Exec(ctx, query, args...)
}

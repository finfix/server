package repository

import (
	"context"
	"fmt"
	"strings"

	"logger/app/logging"
	"pkg/sql"
)

// UpdCurrencies обновляет курсы валют в базе данных
func (repo *Repository) UpdCurrencies(ctx context.Context, rates map[string]float64) error {
	var (
		pattern  = "(?, ?, ?, ?)"
		tmpQuery = make([]string, 0, len(rates))
		args     = make([]interface{}, 0, len(rates)*2)
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

type Repository struct {
	db     *sql.DB
	logger *logging.Logger
}

func New(db *sql.DB, logger *logging.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

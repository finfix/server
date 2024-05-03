package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"

	"server/app/pkg/sql"
	settingsModel "server/app/services/settings/model"
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

func (repo *Repository) GetCurrencies(ctx context.Context) (currencies []settingsModel.Currency, err error) {
	return currencies, repo.db.Select(ctx, &currencies, `SELECT * FROM coin.currencies`)
}

func (repo *Repository) GetIcons(ctx context.Context) (icons []settingsModel.Icon, err error) {
	return icons, repo.db.Select(ctx, &icons, `SELECT * FROM coin.icons`)
}

type Repository struct {
	db sql.SQL
}

func New(db sql.SQL, ) *Repository {
	return &Repository{
		db: db,
	}
}

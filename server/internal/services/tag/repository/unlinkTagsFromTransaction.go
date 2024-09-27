package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// UnlinkTagsFromTransaction отвязывает подкатегории от транзакции
func (r *TagRepository) UnlinkTagsFromTransaction(ctx context.Context, tagIDs []uint32, transactionID uint32) error {

	filtersEq := make(sq.Eq)

	filtersEq["tag_id"] = tagIDs
	filtersEq["transaction_id"] = transactionID

	// Удаляем связь между подкатегориями и транзакцией
	return r.db.Exec(ctx, sq.
		Delete("coin.tags_to_transaction").
		Where(filtersEq),
	)
}

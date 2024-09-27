package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// LinkTagsToTransaction привязывает подкатегории к транзакции
func (r *TagRepository) LinkTagsToTransaction(ctx context.Context, tagIDs []uint32, transactionID uint32) error {

	q := sq.
		Insert("coin.tags_to_transaction").
		Columns("transaction_id", "tag_id")

	for _, tagID := range tagIDs {
		q = q.Values(transactionID, tagID)
	}

	return r.db.Exec(ctx, q)
}

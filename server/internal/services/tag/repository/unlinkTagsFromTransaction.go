package repository

import (
	"context"
	"fmt"
)

// UnlinkTagsFromTransaction отвязывает подкатегории от транзакции
func (repo *TagRepository) UnlinkTagsFromTransaction(ctx context.Context, tagIDs []uint32, transactionID uint32) error {

	args := make([]any, 0, len(tagIDs)+1)

	args = append(args, transactionID)

	queryIn, _args, err := repo.db.In(`tag_id IN (?)`, tagIDs)
	if err != nil {
		return err
	}
	args = append(args, _args...)

	// Удаляем связи между подкатегориями и транзакцией
	query := fmt.Sprintf(`
		DELETE FROM coin.tags_to_transaction
		WHERE transaction_id = ? AND %v`,
		queryIn,
	)

	return repo.db.Exec(ctx, query, args...)
}

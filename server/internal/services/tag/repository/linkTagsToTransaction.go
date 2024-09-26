package repository

import (
	"context"
	"fmt"
	"strings"
)

// LinkTagsToTransaction привязывает подкатегории к транзакции
func (repo *TagRepository) LinkTagsToTransaction(ctx context.Context, tagIDs []uint32, transactionID uint32) error {

	queryValuesTemplate := "(?, ?)"
	queryArgs := make([]string, 0, len(tagIDs))
	args := make([]any, 0, len(tagIDs)*2) // nolint:gomnd

	for _, tagID := range tagIDs {
		queryArgs = append(queryArgs, queryValuesTemplate)
		args = append(args, transactionID, tagID)
	}

	query := fmt.Sprintf(`
		INSERT INTO coin.tags_to_transaction (
            transaction_id,
            tag_id
        ) VALUES %v`,
		strings.Join(queryArgs, ", "),
	)

	return repo.db.Exec(ctx, query, args...)
}

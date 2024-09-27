package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/tag/model"
)

// GetTagsToTransactions возвращает все связи между подкатегориями и транзакциями
func (r *TagRepository) GetTagsToTransactions(ctx context.Context, req model.GetTagsToTransactionsReq) (res []model.TagToTransaction, err error) {

	filtersEq := make(sq.Eq)

	if len(req.AccountGroupIDs) != 0 {
		filtersEq["ag.id"] = req.AccountGroupIDs
	}
	if len(req.TransactionIDs) != 0 {
		filtersEq["ttt.transaction_id"] = req.TransactionIDs
	}

	return res, r.db.Select(ctx, &res, sq.
		Select("*").
		From("coin.tags_to_transaction ttt").
		Join("coin.tags t ON t.id = ttt.tag_id").
		Join("coin.account_groups ag ON ag.id = t.account_group_id").
		Where(filtersEq),
	)
}

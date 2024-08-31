package repository

import (
	"context"
	"fmt"
	"strings"

	"server/internal/services/tag/model"
)

// GetTagsToTransactions возвращает все связи между подкатегориями и транзакциями
func (repo *TagRepository) GetTagsToTransactions(ctx context.Context, req model.GetTagsToTransactionsReq) (res []model.TagToTransaction, err error) {

	var (
		args        []any
		queryFields []string
	)

	if len(req.AccountGroupIDs) != 0 {
		_query, _args, err := repo.db.In(`ag.id IN (?)`, req.AccountGroupIDs)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}
	if len(req.TransactionIDs) != 0 {
		_query, _args, err := repo.db.In(`ttt.transaction_id IN (?)`, req.TransactionIDs)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	request := fmt.Sprintf(`SELECT *
		FROM coin.tags_to_transaction ttt
		JOIN coin.tags t ON t.id = ttt.tag_id
		JOIN coin.account_groups ag ON ag.id = t.account_group_id 
		WHERE %v`,
		strings.Join(queryFields, " AND "),
	)

	return res, repo.db.Select(ctx, &res, request, args...)
}

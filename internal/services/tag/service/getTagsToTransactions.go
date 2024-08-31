package service

import (
	"context"

	"server/internal/services/tag/model"
)

func (s *Service) GetTagsToTransactions(ctx context.Context, req model.GetTagsToTransactionsReq) ([]model.TagToTransaction, error) {

	// Получаем доступные группы счетов
	req.AccountGroupIDs = s.generalRepository.GetAvailableAccountGroups(req.Necessary.UserID)

	// Получаем все связи между подкатегориями и транзакциями
	return s.tagRepository.GetTagsToTransactions(ctx, req)
}

package service

import (
	"context"

	"server/internal/services/generalRepository/checker"
	"server/internal/services/tag/model"
	"server/pkg/slices"
)

func (s *TransactionService) updateTransactionTags(ctx context.Context, userID, transactionID uint32, tagIDs []uint32) error {

	// Проверяем доступ пользователя к тегам
	if len(tagIDs) > 0 {
		if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Tags, userID, tagIDs); err != nil {
			return err
		}
	}

	// Получаем все теги, привязанные к транзакции
	transactionTags, err := s.tagRepository.GetTagsToTransactions(ctx, model.GetTagsToTransactionsReq{ //nolint:exhaustruct
		TransactionIDs: []uint32{transactionID},
	})
	if err != nil {
		return err
	}

	existTagIDs := slices.GetFields(transactionTags, func(tag model.TagToTransaction) uint32 { return tag.TagID })

	toDelete, toCreate := slices.JoinExclusive(existTagIDs, tagIDs)

	if len(toDelete) > 0 {
		// Удаляем связи тегов с транзакцией
		if err = s.tagRepository.UnlinkTagsFromTransaction(ctx, toDelete, transactionID); err != nil {
			return err
		}
	}
	if len(toCreate) > 0 {
		// Создаем связи тегов с транзакцией
		if err = s.tagRepository.LinkTagsToTransaction(ctx, toCreate, transactionID); err != nil {
			return err
		}
	}
	return nil
}

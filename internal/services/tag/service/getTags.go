package service

import (
	"context"

	"server/internal/services/generalRepository/checker"
	"server/internal/services/tag/model"
)

func (s *Service) GetTags(ctx context.Context, filters model.GetTagsReq) (tags []model.Tag, err error) {

	// Проверяем доступ пользователя к группам счетов
	if filters.AccountGroupIDs != nil {
		if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.AccountGroups, filters.Necessary.UserID, filters.AccountGroupIDs); err != nil {
			return nil, err
		}
	} else {
		filters.AccountGroupIDs = s.generalRepository.GetAvailableAccountGroups(filters.Necessary.UserID)
	}

	// Получаем все подкатегории
	if tags, err = s.tagRepository.GetTags(ctx, filters); err != nil {
		return nil, err
	}

	// Заполняем массив ID транзакций
	tagIDs := make([]uint32, len(tags))
	for i, tag := range tags {
		tagIDs[i] = tag.ID
	}

	return tags, nil
}

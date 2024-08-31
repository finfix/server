package service

import (
	"context"

	"server/internal/services/generalRepository/checker"
	"server/internal/services/tag/model"
)

// CreateTag создает новую подкатегорию
func (s *Service) CreateTag(ctx context.Context, tag model.CreateTagReq) (id uint32, err error) {

	// Проверяем доступ пользователя к счетам
	if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.AccountGroups, tag.Necessary.UserID, []uint32{tag.AccountGroupID}); err != nil {
		return id, err
	}

	// Создаем подкатегорию
	return s.tagRepository.CreateTag(ctx, tag.ConvertToRepoReq())
}

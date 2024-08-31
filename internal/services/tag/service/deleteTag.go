package service

import (
	"context"

	"server/internal/services/generalRepository/checker"
	"server/internal/services/tag/model"
)

// DeleteTag удаляет подкатегорию
func (s *Service) DeleteTag(ctx context.Context, id model.DeleteTagReq) error {

	// Проверяем доступ пользователя к подкатегории
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Tags, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем подкатегорию
	return s.tagRepository.DeleteTag(ctx, id.ID, id.Necessary.UserID)
}

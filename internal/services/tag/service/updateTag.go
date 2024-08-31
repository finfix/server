package service

import (
	"context"

	"server/internal/services/generalRepository/checker"
	"server/internal/services/tag/model"
)

// UpdateTag редактирует подкатегорию
func (s *Service) UpdateTag(ctx context.Context, fields model.UpdateTagReq) error {

	// Проверяем доступ пользователя к подкатегории
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Tags, fields.Necessary.UserID, []uint32{fields.ID}); err != nil {
		return err
	}

	// Изменяем данные подкатегории
	return s.tagRepository.UpdateTag(ctx, fields)
}

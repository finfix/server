package service

import (
	"context"

	"server/app/services/generalRepository"
	"server/app/services/generalRepository/checker"
	tagModel "server/app/services/tag/model"
	tagRepository "server/app/services/tag/repository"
	tagRepoModel "server/app/services/tag/repository/model"
)

type Service struct {
	tagRepository     TagRepository
	generalRepository GeneralRepository
}

var _ TagRepository = &tagRepository.TagRepository{}
var _ GeneralRepository = &generalRepository.Repository{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	CheckUserAccessToObjects(context.Context, checker.CheckType, uint32, []uint32) error
	GetAvailableAccountGroups(userID uint32) []uint32
}

type TagRepository interface {
	CreateTag(context.Context, tagRepoModel.CreateTagReq) (uint32, error)
	UpdateTag(context.Context, tagModel.UpdateTagReq) error
	DeleteTag(ctx context.Context, id, userID uint32) error
	GetTags(context.Context, tagModel.GetTagsReq) (res []tagModel.Tag, err error)

	GetTagsToTransactions(ctx context.Context, req tagModel.GetTagsToTransactionsReq) ([]tagModel.TagToTransaction, error)
}

// CreateTag создает новую подкатегорию
func (s *Service) CreateTag(ctx context.Context, tag tagModel.CreateTagReq) (id uint32, err error) {

	// Проверяем доступ пользователя к счетам
	if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.AccountGroups, tag.Necessary.UserID, []uint32{tag.AccountGroupID}); err != nil {
		return id, err
	}

	// Создаем подкатегорию
	return s.tagRepository.CreateTag(ctx, tag.ConvertToRepoReq())
}

func (s *Service) GetTags(ctx context.Context, filters tagModel.GetTagsReq) (tags []tagModel.Tag, err error) {

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

// UpdateTag редактирует подкатегорию
func (s *Service) UpdateTag(ctx context.Context, fields tagModel.UpdateTagReq) error {

	// Проверяем доступ пользователя к подкатегории
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Tags, fields.Necessary.UserID, []uint32{fields.ID}); err != nil {
		return err
	}

	// Изменяем данные подкатегории
	return s.tagRepository.UpdateTag(ctx, fields)
}

// DeleteTag удаляет подкатегорию
func (s *Service) DeleteTag(ctx context.Context, id tagModel.DeleteTagReq) error {

	// Проверяем доступ пользователя к подкатегории
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Tags, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем подкатегорию
	return s.tagRepository.DeleteTag(ctx, id.ID, id.Necessary.UserID)
}

func (s *Service) GetTagsToTransactions(ctx context.Context, req tagModel.GetTagsToTransactionsReq) ([]tagModel.TagToTransaction, error) {

	// Получаем доступные группы счетов
	req.AccountGroupIDs = s.generalRepository.GetAvailableAccountGroups(req.Necessary.UserID)

	// Получаем все связи между подкатегориями и транзакциями
	return s.tagRepository.GetTagsToTransactions(ctx, req)
}

func New(
	tagRepository TagRepository,
	generalRepository GeneralRepository,

) *Service {
	return &Service{
		tagRepository:     tagRepository,
		generalRepository: generalRepository,
	}
}

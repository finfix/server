package service

import (
	"context"

	"server/internal/services/generalRepository"
	"server/internal/services/generalRepository/checker"
	tagModel "server/internal/services/tag/model"
	tagRepository "server/internal/services/tag/repository"
	tagRepoModel "server/internal/services/tag/repository/model"
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

func New(
	tagRepository TagRepository,
	generalRepository GeneralRepository,

) *Service {
	return &Service{
		tagRepository:     tagRepository,
		generalRepository: generalRepository,
	}
}

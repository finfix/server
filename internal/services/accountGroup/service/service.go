package service

import (
	"context"

	"github.com/shopspring/decimal"

	accountGroupModel "server/internal/services/accountGroup/model"
	accountGroupRepository "server/internal/services/accountGroup/repository"
	accountGroupRepoModel "server/internal/services/accountGroup/repository/model"
	"server/internal/services/generalRepository"
	"server/internal/services/generalRepository/checker"
)

var _ GeneralRepository = &generalRepository.Repository{}
var _ AccountGroupRepository = &accountGroupRepository.Repository{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	GetCurrencies(context.Context) (map[string]decimal.Decimal, error)
	CheckUserAccessToObjects(context.Context, checker.CheckType, uint32, []uint32) error
}

type AccountGroupRepository interface {
	CreateAccountGroup(context.Context, accountGroupRepoModel.CreateAccountGroupReq) (uint32, uint32, error)
	GetAccountGroups(context.Context, accountGroupModel.GetAccountGroupsReq) ([]accountGroupModel.AccountGroup, error)
	UpdateAccountGroup(context.Context, accountGroupModel.UpdateAccountGroupReq) error
	DeleteAccountGroup(ctx context.Context, id uint32) error

	LinkUserToAccountGroup(ctx context.Context, userID, accountGroupID uint32) error
	UnlinkUserFromAccountGroup(ctx context.Context, userID, accountGroupID uint32) error
}

type Service struct {
	accountGroupRepository AccountGroupRepository
	general                GeneralRepository
}

func New(
	accountGroup AccountGroupRepository,
	general GeneralRepository,

) *Service {
	return &Service{
		accountGroupRepository: accountGroup,
		general:                general,
	}
}

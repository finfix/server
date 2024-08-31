package service

import (
	"context"

	accountModel "server/internal/services/account/model"
	accountRepository "server/internal/services/account/repository"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/internal/services/accountPermissions"
	"server/internal/services/generalRepository"
	"server/internal/services/generalRepository/checker"
	tagModel "server/internal/services/tag/model"
	tagRepository "server/internal/services/tag/repository"
	transactionModel "server/internal/services/transaction/model"
	transactionRepository "server/internal/services/transaction/repository"
	transactionRepoModel "server/internal/services/transaction/repository/model"
)

type Service struct {
	transactionRepository TransactionRepository
	accountRepository     AccountRepository
	generalRepository     GeneralRepository
	permissionsService    AccountPermissionsService
	tagRepository         TagRepository
}

var _ TransactionRepository = &transactionRepository.TransactionRepository{}
var _ AccountRepository = &accountRepository.Repository{}
var _ GeneralRepository = &generalRepository.Repository{}
var _ AccountPermissionsService = &accountPermissions.Service{}
var _ TagRepository = &tagRepository.TagRepository{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	CheckUserAccessToObjects(context.Context, checker.CheckType, uint32, []uint32) error
	GetAvailableAccountGroups(userID uint32) []uint32
}

type TransactionRepository interface {
	CreateTransaction(context.Context, transactionRepoModel.CreateTransactionReq) (uint32, error)
	UpdateTransaction(context.Context, transactionModel.UpdateTransactionReq) error
	DeleteTransaction(ctx context.Context, id, userID uint32) error
	GetTransactions(context.Context, transactionModel.GetTransactionsReq) (res []transactionModel.Transaction, err error)
}

type AccountPermissionsService interface {
	GetAccountPermissions(account accountModel.Account) accountPermissions.AccountPermissions
}

type AccountRepository interface {
	GetAccounts(context.Context, accountRepoModel.GetAccountsReq) ([]accountModel.Account, error)
}

type TagRepository interface {
	GetTagsToTransactions(context.Context, tagModel.GetTagsToTransactionsReq) ([]tagModel.TagToTransaction, error)
	LinkTagsToTransaction(context.Context, []uint32, uint32) error
	UnlinkTagsFromTransaction(context.Context, []uint32, uint32) error
}

func New(
	transactionRepository TransactionRepository,
	accountRepository AccountRepository,
	generalRepository GeneralRepository,
	accountPermissions AccountPermissionsService,
	tagRepository TagRepository,

) *Service {
	return &Service{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
		generalRepository:     generalRepository,
		permissionsService:    accountPermissions,
		tagRepository:         tagRepository,
	}
}

package service

import (
	"context"

	accountModel "server/internal/services/account/model"
	accountRepository "server/internal/services/account/repository"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/internal/services/accountPermissions/model"
	"server/internal/services/accountPermissions/service"
	"server/internal/services/generalRepository"
	"server/internal/services/generalRepository/checker"
	tagModel "server/internal/services/tag/model"
	tagRepository "server/internal/services/tag/repository"
	transactionModel "server/internal/services/transaction/model"
	transactionRepository "server/internal/services/transaction/repository"
	transactionRepoModel "server/internal/services/transaction/repository/model"
)

type TransactionService struct {
	transactionRepository TransactionRepository
	accountRepository     AccountRepository
	generalRepository     GeneralRepository
	permissionsService    AccountPermissionsService
	tagRepository         TagRepository
}

var _ TransactionRepository = new(transactionRepository.TransactionRepository)
var _ AccountRepository = new(accountRepository.AccountRepository)
var _ GeneralRepository = new(generalRepository.GeneralRepository)
var _ AccountPermissionsService = new(service.AccountPermissionsService)
var _ TagRepository = new(tagRepository.TagRepository)

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
	GetAccountsPermissions(context.Context, ...accountModel.Account) ([]model.AccountPermissions, error)
}

type AccountRepository interface {
	GetAccounts(context.Context, accountRepoModel.GetAccountsReq) ([]accountModel.Account, error)
}

type TagRepository interface {
	GetTagsToTransactions(context.Context, tagModel.GetTagsToTransactionsReq) ([]tagModel.TagToTransaction, error)
	LinkTagsToTransaction(context.Context, []uint32, uint32) error
	UnlinkTagsFromTransaction(context.Context, []uint32, uint32) error
}

func NewTransactionService(
	transactionRepository TransactionRepository,
	accountRepository AccountRepository,
	generalRepository GeneralRepository,
	accountPermissions AccountPermissionsService,
	tagRepository TagRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
		generalRepository:     generalRepository,
		permissionsService:    accountPermissions,
		tagRepository:         tagRepository,
	}
}

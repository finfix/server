package service

import (
	"context"

	"github.com/shopspring/decimal"

	accountModel "server/internal/services/account/model"
	accountRepository "server/internal/services/account/repository"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/internal/services/accountPermissions/model"
	"server/internal/services/accountPermissions/service"
	"server/internal/services/generalRepository"
	"server/internal/services/generalRepository/checker"
	transactionRepository "server/internal/services/transaction/repository"
	transactionRepoModel "server/internal/services/transaction/repository/model"
	userModel "server/internal/services/user/model"
	userRepository "server/internal/services/user/repository"
)

var _ GeneralRepository = new(generalRepository.GeneralRepository)
var _ AccountRepository = new(accountRepository.AccountRepository)
var _ AccountPermissionsService = new(service.AccountPermissionsService)
var _ UserRepository = new(userRepository.UserRepository)
var _ TransactionRepository = new(transactionRepository.TransactionRepository)

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	GetCurrencies(context.Context) (map[string]decimal.Decimal, error)
	CheckUserAccessToObjects(context.Context, checker.CheckType, uint32, []uint32) error
	GetAvailableAccountGroups(userID uint32) []uint32
}

type AccountRepository interface {
	CreateAccount(context.Context, accountRepoModel.CreateAccountReq) (uint32, uint32, error)
	GetAccounts(context.Context, accountRepoModel.GetAccountsReq) ([]accountModel.Account, error)
	UpdateAccount(context.Context, map[uint32]accountRepoModel.UpdateAccountReq) error
	DeleteAccount(ctx context.Context, id uint32) error

	ChangeSerialNumbers(ctx context.Context, accountGroupID, oldValue, newValue uint32) error
	CalculateRemainderAccounts(ctx context.Context, req accountRepoModel.CalculateRemaindersAccountsReq) (map[uint32]decimal.Decimal, error)
}

type TransactionRepository interface {
	CreateTransaction(context.Context, transactionRepoModel.CreateTransactionReq) (uint32, error)
}

type UserRepository interface {
	GetUsers(context.Context, userModel.GetUsersReq) ([]userModel.User, error)
}

type AccountPermissionsService interface {
	GetAccountsPermissions(context.Context, ...accountModel.Account) ([]model.AccountPermissions, error)
}

type AccountService struct {
	accountRepository         AccountRepository
	general                   GeneralRepository
	transaction               TransactionRepository
	user                      UserRepository
	accountPermissionsService AccountPermissionsService
}

func NewAccountService(
	account AccountRepository,
	general GeneralRepository,
	transaction TransactionRepository,
	user UserRepository,
	permissionsService AccountPermissionsService,
) *AccountService {
	return &AccountService{
		accountRepository:         account,
		general:                   general,
		transaction:               transaction,
		user:                      user,
		accountPermissionsService: permissionsService,
	}
}

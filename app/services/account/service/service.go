package service

import (
	"context"

	"github.com/shopspring/decimal"

	accountModel "server/app/services/account/model"
	accountRepository "server/app/services/account/repository"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/accountPermissions"
	"server/app/services/generalRepository"
	"server/app/services/generalRepository/checker"
	transactionRepository "server/app/services/transaction/repository"
	transactionRepoModel "server/app/services/transaction/repository/model"
	userModel "server/app/services/user/model"
	userRepository "server/app/services/user/repository"
)

var _ GeneralRepository = &generalRepository.Repository{}
var _ AccountRepository = &accountRepository.Repository{}
var _ AccountService = &Service{} //nolint:exhaustruct
var _ AccountPermissionsService = &accountPermissions.Service{}
var _ UserRepository = &userRepository.Repository{}
var _ TransactionRepository = &transactionRepository.TransactionRepository{}

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
	GetUsers(context.Context, userModel.GetReq) ([]userModel.User, error)
}

type AccountPermissionsService interface {
	GetAccountPermissions(account accountModel.Account) accountPermissions.AccountPermissions
	CheckAccountPermissions(req accountModel.UpdateAccountReq, permissions accountPermissions.AccountPermissions) error
}

type AccountService interface {
	ChangeAccountRemainder(ctx context.Context, account accountModel.Account, remainderToUpdate decimal.Decimal, userID uint32) (accountModel.UpdateAccountRes, error)
}

type Service struct {
	accountService            AccountService
	accountRepository         AccountRepository
	general                   GeneralRepository
	transaction               TransactionRepository
	user                      UserRepository
	accountPermissionsService AccountPermissionsService
}

func New(
	account AccountRepository,
	general GeneralRepository,
	transaction TransactionRepository,
	user UserRepository,
	permissionsService AccountPermissionsService,

) *Service {
	s := &Service{
		accountRepository:         account,
		general:                   general,
		transaction:               transaction,
		user:                      user,
		accountPermissionsService: permissionsService,
		accountService:            nil,
	}
	s.accountService = s
	return s
}

package service

import (
	"context"

	"server/app/pkg/logging"
	accountModel "server/app/services/account/model"
	accountRepository "server/app/services/account/repository"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository"
	"server/app/services/generalRepository/checker"
	"server/app/services/permissions"
	transactionModel "server/app/services/transaction/model"
	transactionRepository "server/app/services/transaction/repository"
	userModel "server/app/services/user/model"
	userRepository "server/app/services/user/repository"
)

var _ GeneralRepository = &generalRepository.Repository{}
var _ AccountRepository = &accountRepository.Repository{}
var _ PermissionsService = &permissions.Service{}
var _ UserRepository = &userRepository.Repository{}
var _ TransactionRepository = &transactionRepository.TransactionRepository{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	GetCurrencies(context.Context) (map[string]float64, error)
	CheckAccess(context.Context, checker.CheckType, uint32, []uint32) error
	GetAvailableAccountGroups(userID uint32) []uint32
}

type AccountRepository interface {
	Create(context.Context, accountRepoModel.CreateReq) (uint32, uint32, error)
	Get(context.Context, accountRepoModel.GetReq) ([]accountModel.Account, error)
	Update(context.Context, map[uint32]accountRepoModel.UpdateReq) error
	Delete(ctx context.Context, id uint32) error

	CalculateRemainderAccounts(ctx context.Context, req accountRepoModel.CalculateRemaindersAccountsReq) (map[uint32]float64, error)
	Switch(ctx context.Context, id1, id2 uint32) error

	GetAccountGroups(context.Context, accountModel.GetAccountGroupsReq) ([]accountModel.AccountGroup, error)
	CreateAccountGroup(context.Context, accountModel.CreateAccountGroupReq) (uint32, error)
}

type TransactionRepository interface {
	Create(context.Context, transactionModel.CreateReq) (uint32, error)
}

type UserRepository interface {
	Get(context.Context, userModel.GetReq) ([]userModel.User, error)
}

type PermissionsService interface {
	GetPermissions(account accountModel.Account) permissions.Permissions
	CheckPermissions(req accountModel.UpdateReq, permissions permissions.Permissions) error
}

type AccountService interface {
	ChangeRemainder(ctx context.Context, account accountModel.Account, remainderToUpdate float64) (accountModel.UpdateRes, error)
	ValidateUpdateParentAccountID(ctx context.Context, account accountModel.Account, parentAccountID, userID uint32) error
}

type Service struct {
	accountService     AccountService
	accountRepository  AccountRepository
	general            GeneralRepository
	transaction        TransactionRepository
	user               UserRepository
	permissionsService PermissionsService
	logger             *logging.Logger
}

func New(
	account AccountRepository,
	general GeneralRepository,
	transaction TransactionRepository,
	user UserRepository,
	permissionsService PermissionsService,
	logger *logging.Logger,
) *Service {
	s := &Service{
		accountRepository:  account,
		general:            general,
		transaction:        transaction,
		user:               user,
		permissionsService: permissionsService,
		logger:             logger,
	}
	s.accountService = s
	return s
}

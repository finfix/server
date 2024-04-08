package service

import (
	"context"

	model3 "server/app/services/account/model"
	accountRepository "server/app/services/account/repository"
	"server/app/services/generalRepository"
	"server/app/services/generalRepository/checker"
	"server/app/services/permissions"
	transactionModel "server/app/services/transaction/model"
	transactionRepository "server/app/services/transaction/repository"
	model2 "server/app/services/user/model"
	userRepository "server/app/services/user/repository"
	"server/pkg/datetime/date"
	"server/pkg/logging"
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
	Create(context.Context, model3.CreateReq) (uint32, uint32, error)
	Get(context.Context, model3.GetReq) ([]model3.Account, error)
	Update(context.Context, model3.UpdateReq) error
	Delete(ctx context.Context, id uint32) error

	GetRemainder(ctx context.Context, id uint32) (float64, error)
	CalculateExpensesAndEarnings(ctx context.Context, accountGroupIDs []uint32, dateFrom, dateTo date.Date) (map[uint32]float64, error)
	CalculateRemainderAccounts(ctx context.Context, accountGroupIDs []uint32, dateTo *date.Date) (map[uint32]float64, error)
	CalculateBalancingAmount(ctx context.Context, accountGroupIDs []uint32, dateFrom, dateTo date.Date) ([]model3.BalancingAmount, error)
	Switch(ctx context.Context, id1, id2 uint32) error

	GetAccountGroups(context.Context, model3.GetAccountGroupsReq) ([]model3.AccountGroup, error)
	CreateAccountGroup(context.Context, model3.CreateAccountGroupReq) (uint32, error)
}

type TransactionRepository interface {
	Create(context.Context, transactionModel.CreateReq) (uint32, error)
}

type UserRepository interface {
	Get(context.Context, model2.GetReq) ([]model2.User, error)
}

type PermissionsService interface {
	GetPermissions(account model3.Account) permissions.Permissions
	CheckPermissions(req model3.UpdateReq, permissions permissions.Permissions) error
}

type Service struct {
	account            AccountRepository
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
	return &Service{
		account:            account,
		general:            general,
		transaction:        transaction,
		user:               user,
		permissionsService: permissionsService,
		logger:             logger,
	}
}

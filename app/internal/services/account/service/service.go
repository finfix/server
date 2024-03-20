package service

import (
	"context"
	"server/app/internal/services/generalRepository"
	"server/app/internal/services/generalRepository/checker"

	"server/pkg/datetime/date"
	"server/pkg/logging"

	"server/app/internal/services/account/model"
	accountRepository "server/app/internal/services/account/repository"
	transactionModel "server/app/internal/services/transaction/model"
	userModel "server/app/internal/services/user/model"
)

var _ GeneralRepository = &generalRepository.Repository{}
var _ AccountRepository = &accountRepository.Repository{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	GetCurrencies(context.Context) (map[string]float64, error)
	CheckAccess(context.Context, checker.CheckType, uint32, []uint32) error
	GetAvailableAccountGroups(userID uint32) []uint32
}

type AccountRepository interface {
	Create(context.Context, model.CreateReq) (uint32, error)
	Get(context.Context, model.GetReq) ([]model.Account, error)
	Update(context.Context, model.UpdateReq) error
	Delete(ctx context.Context, id uint32) error

	GetRemainder(ctx context.Context, id uint32) (float64, error)
	CalculateExpensesAndEarnings(ctx context.Context, accountGroupIDs []uint32, dateFrom, dateTo date.Date) (map[uint32]float64, error)
	CalculateRemainderAccounts(ctx context.Context, accountGroupIDs []uint32, dateTo *date.Date) (map[uint32]float64, error)
	CalculateBalancingAmount(ctx context.Context, accountGroupIDs []uint32, dateFrom, dateTo date.Date) ([]model.BalancingAmount, error)
	Switch(ctx context.Context, id1, id2 uint32) error

	GetAccountGroups(context.Context, model.GetAccountGroupsReq) ([]model.AccountGroup, error)
	CreateAccountGroup(context.Context, model.CreateAccountGroupReq) (uint32, error)
}

type TransactionRepository interface {
	Create(context.Context, transactionModel.CreateReq) (uint32, error)
}

type UserRepository interface {
	Get(context.Context, userModel.GetReq) ([]userModel.User, error)
}

type Service struct {
	account     AccountRepository
	general     GeneralRepository
	transaction TransactionRepository
	user        UserRepository
	logger      *logging.Logger
}

func New(
	account AccountRepository,
	general GeneralRepository,
	transaction TransactionRepository,
	user UserRepository,
	logger *logging.Logger,
) *Service {
	return &Service{
		account:     account,
		general:     general,
		transaction: transaction,
		user:        user,
		logger:      logger,
	}
}

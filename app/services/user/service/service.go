package service

import (
	"context"

	"server/app/pkg/logging"
	accountModel "server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository"
	model2 "server/app/services/user/model"
	userRepository "server/app/services/user/repository"
)

var _ UserRepository = &userRepository.Repository{}
var _ GeneralRepository = &generalRepository.Repository{}

type UserRepository interface {
	Create(context.Context, model2.CreateReq) (uint32, error)
	GetTransactions(context.Context, model2.GetReq) ([]model2.User, error)
	GetCurrencies(context.Context) ([]model2.Currency, error)
	LinkUserToAccountGroup(context.Context, uint32, uint32) error
}

type AccountRepository interface {
	CreateAccountGroup(context.Context, accountModel.CreateAccountGroupReq) (uint32, error)
	CreateAccount(ctx context.Context, req accountRepoModel.CreateAccountReq) (uint32, uint32, error)
}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
}

type Service struct {
	user    UserRepository
	account AccountRepository
	general GeneralRepository
	logger  *logging.Logger
}

// Create создает нового пользователя
func (s *Service) Create(ctx context.Context, user model2.CreateReq) (id uint32, err error) {
	return s.user.Create(ctx, user)
}

// GetTransactions возвращает всех юзеров по фильтрам
func (s *Service) GetTransactions(ctx context.Context, filters model2.GetReq) (users []model2.User, err error) {
	return s.user.GetTransactions(ctx, filters)
}

func (s *Service) GetCurrencies(ctx context.Context) ([]model2.Currency, error) {
	return s.user.GetCurrencies(ctx)
}

func New(
	user UserRepository,
	general GeneralRepository,
	account AccountRepository,
	logger *logging.Logger,
) *Service {
	return &Service{
		user:    user,
		general: general,
		account: account,
		logger:  logger,
	}
}

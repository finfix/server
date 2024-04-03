package service

import (
	"context"

	accountModel "server/app/internal/services/account/model"
	"server/app/internal/services/generalRepository"
	"server/app/internal/services/user/model"
	userRepository "server/app/internal/services/user/repository"
	"server/pkg/logging"
)

var _ UserRepository = &userRepository.Repository{}
var _ GeneralRepository = &generalRepository.Repository{}

type UserRepository interface {
	Create(context.Context, model.CreateReq) (uint32, error)
	Get(context.Context, model.GetReq) ([]model.User, error)
	GetCurrencies(context.Context) ([]model.Currency, error)
	LinkUserToAccountGroup(context.Context, uint32, uint32) error
}

type AccountRepository interface {
	CreateAccountGroup(context.Context, accountModel.CreateAccountGroupReq) (uint32, error)
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
func (s *Service) Create(ctx context.Context, user model.CreateReq) (id uint32, err error) {

	err = s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем пользователя
		if id, err = s.user.Create(ctx, user); err != nil {
			return err
		}

		// Создаем дефолтную группу счетов с новой группой юзеров
		accountGroupID, err := s.account.CreateAccountGroup(ctx, accountModel.CreateAccountGroupReq{
			Name:            "Личные",
			AvailableBudget: 0,     // TODO: Передавать в запросе
			Currency:        "RUB", // TODO: Передавать в запросе
		})
		if err != nil {
			return err
		}

		// Связываем юзера с группой юзеров
		return s.user.LinkUserToAccountGroup(ctx, id, accountGroupID)
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get возвращает всех юзеров по фильтрам
func (s *Service) Get(ctx context.Context, filters model.GetReq) (users []model.User, err error) {
	return s.user.Get(ctx, filters)
}

func (s *Service) GetCurrencies(ctx context.Context) ([]model.Currency, error) {
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

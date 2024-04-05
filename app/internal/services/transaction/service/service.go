package service

import (
	"context"

	accountService "server/app/internal/services/account/service"
	"server/app/internal/services/generalRepository"
	"server/app/internal/services/generalRepository/checker"
	"server/app/internal/services/transaction/model"
	transactionRepository "server/app/internal/services/transaction/repository"
	"server/pkg/errors"
	"server/pkg/logging"
)

var _ GeneralRepository = &generalRepository.Repository{}
var _ Repository = &transactionRepository.TransactionRepository{}
var _ AccountService = &accountService.Service{}

type AccountService interface {
	GetPermissions(ctx context.Context, id uint32) (accountService.Permissions, error)
}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	CheckAccess(context.Context, checker.CheckType, uint32, []uint32) error
	GetAvailableAccountGroups(userID uint32) []uint32
}

type Repository interface {
	Create(context.Context, model.CreateReq) (uint32, error)
	Update(context.Context, model.UpdateReq) error
	Delete(ctx context.Context, id, userID uint32) error
	Get(context.Context, model.GetReq) (res []model.Transaction, err error)

	CreateTags(ctx context.Context, tags []string, transactionID uint32) error
	GetTags(ctx context.Context, transactionID []uint32) ([]model.Tag, error)
}

// Create создает новую транзакцию
func (s *Service) Create(ctx context.Context, transaction model.CreateReq) (id uint32, err error) {

	// Проверяем доступ пользователя к счетам
	if err = s.general.CheckAccess(ctx, checker.Accounts, transaction.UserID, []uint32{transaction.AccountFromID, transaction.AccountToID}); err != nil {
		return id, err
	}

	// Получаем разрешения счетов
	permissions1, err := s.account.GetPermissions(ctx, transaction.AccountFromID)
	if err != nil {
		return id, err
	}
	permissions2, err := s.account.GetPermissions(ctx, transaction.AccountToID)
	if err != nil {
		return id, err
	}

	// Проверяем, что счета можно использовать
	if !permissions1.CreateTransaction || !permissions2.CreateTransaction {
		return id, errors.BadRequest.New("Нельзя создать транзакцию для этих счетов")
	}

	// Создаем транзакцию
	return s.transaction.Create(ctx, transaction)
}

func (s *Service) Get(ctx context.Context, filters model.GetReq) (transactions []model.Transaction, err error) {

	// Проверяем доступ пользователя к группам счетов
	filters.AccountGroupIDs = s.general.GetAvailableAccountGroups(filters.UserID)

	// Получаем все транзакции
	if transactions, err = s.transaction.Get(ctx, filters); err != nil {
		return nil, err
	}

	// Заполняем массив ID транзакций
	transactionIDs := make([]uint32, len(transactions))
	for i, transaction := range transactions {
		transactionIDs[i] = transaction.ID
	}

	return transactions, nil
}

// Update редактирует транзакцию
func (s *Service) Update(ctx context.Context, fields model.UpdateReq) error {

	// Проверяем доступ пользователя к транзакции
	if err := s.general.CheckAccess(ctx, checker.Transactions, fields.UserID, []uint32{fields.ID}); err != nil {
		return err
	}

	// Если в запросе есть изменение счетов, то проверяем доступ пользователя к ним
	if fields.AccountFromID != nil || fields.AccountToID != nil {
		var accountsIDs []uint32
		if fields.AccountFromID != nil {
			accountsIDs = append(accountsIDs, *fields.AccountFromID)
		}
		if fields.AccountToID != nil {
			accountsIDs = append(accountsIDs, *fields.AccountToID)
		}

		// Проверяем доступ пользователя к счетам
		if err := s.general.CheckAccess(ctx, checker.Accounts, fields.UserID, accountsIDs); err != nil {
			return err
		}
	}

	// Проверяем доступ пользователя к транзакции
	if err := s.general.CheckAccess(ctx, checker.Transactions, fields.UserID, []uint32{fields.ID}); err != nil {
		return err
	}

	// Изменяем данные транзакции
	return s.transaction.Update(ctx, fields)
}

// Delete удаляет транзакцию
func (s *Service) Delete(ctx context.Context, id model.DeleteReq) error {

	// Проверяем доступ пользователя к транзакции
	if err := s.general.CheckAccess(ctx, checker.Transactions, id.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем транзакцию
	return s.transaction.Delete(ctx, id.ID, id.UserID)
}

type Service struct {
	transaction Repository
	account     AccountService
	general     GeneralRepository
	logger      *logging.Logger
}

func New(
	transactionRepository Repository,
	accountService AccountService,
	generalRepository GeneralRepository,
	logger *logging.Logger,
) *Service {
	return &Service{
		transaction: transactionRepository,
		account:     accountService,
		general:     generalRepository,
		logger:      logger,
	}
}

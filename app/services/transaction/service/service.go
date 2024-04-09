package service

import (
	"context"

	model3 "server/app/services/account/model"
	"server/app/services/generalRepository"
	"server/app/services/generalRepository/checker"
	"server/app/services/permissions"
	model2 "server/app/services/transaction/model"
	transactionRepository "server/app/services/transaction/repository"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/slice"
)

type Service struct {
	transaction Repository
	account     AccountRepository
	general     GeneralRepository
	permissions PermissionsService
	logger      *logging.Logger
}

var _ GeneralRepository = &generalRepository.Repository{}
var _ Repository = &transactionRepository.TransactionRepository{}
var _ PermissionsService = &permissions.Service{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
	CheckAccess(context.Context, checker.CheckType, uint32, []uint32) error
	GetAvailableAccountGroups(userID uint32) []uint32
}

type Repository interface {
	Create(context.Context, model2.CreateReq) (uint32, error)
	Update(context.Context, model2.UpdateReq) error
	Delete(ctx context.Context, id, userID uint32) error
	Get(context.Context, model2.GetReq) (res []model2.Transaction, err error)

	CreateTags(ctx context.Context, tags []string, transactionID uint32) error
	GetTags(ctx context.Context, transactionID []uint32) ([]model2.Tag, error)
}

type PermissionsService interface {
	GetPermissions(account model3.Account) permissions.Permissions
}

type AccountRepository interface {
	Get(context.Context, model3.GetReq) ([]model3.Account, error)
}

// Create создает новую транзакцию
func (s *Service) Create(ctx context.Context, transaction model2.CreateReq) (id uint32, err error) {

	// Проверяем доступ пользователя к счетам
	if err = s.general.CheckAccess(ctx, checker.Accounts, transaction.UserID, []uint32{transaction.AccountFromID, transaction.AccountToID}); err != nil {
		return id, err
	}

	// Получаем счета
	_accounts, err := s.account.Get(ctx, model3.GetReq{
		IDs: []uint32{transaction.AccountFromID, transaction.AccountToID},
	})
	if err != nil {
		return id, err
	}
	accountsMap := slice.ToMap(_accounts, func(account model3.Account) uint32 { return account.ID })

	// Получаем разрешения счетов
	permissionsAccountFrom := s.permissions.GetPermissions(accountsMap[transaction.AccountFromID])
	permissionsAccountTo := s.permissions.GetPermissions(accountsMap[transaction.AccountToID])

	// Проверяем, что счета можно использовать
	if !permissionsAccountFrom.CreateTransaction || !permissionsAccountTo.CreateTransaction {
		return id, errors.BadRequest.New("Нельзя создать транзакцию для этих счетов", errors.Options{
			Params: map[string]any{
				"AccountFromID":      transaction.AccountFromID,
				"AccountGroupFromID": accountsMap[transaction.AccountFromID].AccountGroupID,
				"AccountToID":        transaction.AccountToID,
				"AccountGroupToID":   accountsMap[transaction.AccountToID].AccountGroupID,
			},
		})
	}

	// Проверяем, что счета находятся в одной группе
	if accountsMap[transaction.AccountFromID].AccountGroupID != accountsMap[transaction.AccountToID].AccountGroupID {
		return id, errors.BadRequest.New("Счета находятся в разных группах", errors.Options{
			Params: map[string]any{
				"AccountFromID": transaction.AccountFromID,
				"AccountToID":   transaction.AccountToID,
			},
		})
	}

	// Создаем транзакцию
	return s.transaction.Create(ctx, transaction)
}

func (s *Service) Get(ctx context.Context, filters model2.GetReq) (transactions []model2.Transaction, err error) {

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
func (s *Service) Update(ctx context.Context, fields model2.UpdateReq) error {

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
func (s *Service) Delete(ctx context.Context, id model2.DeleteReq) error {

	// Проверяем доступ пользователя к транзакции
	if err := s.general.CheckAccess(ctx, checker.Transactions, id.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем транзакцию
	return s.transaction.Delete(ctx, id.ID, id.UserID)
}

func New(
	transactionRepository Repository,
	accountRepository AccountRepository,
	generalRepository GeneralRepository,
	permissions PermissionsService,
	logger *logging.Logger,
) *Service {
	return &Service{
		transaction: transactionRepository,
		account:     accountRepository,
		general:     generalRepository,
		permissions: permissions,
		logger:      logger,
	}
}

package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/slice"
	accountModel "server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/accountPermissions"
	"server/app/services/generalRepository"
	"server/app/services/generalRepository/checker"
	transactionModel "server/app/services/transaction/model"
	transactionRepository "server/app/services/transaction/repository"
	transactionRepoModel "server/app/services/transaction/repository/model"
)

type Service struct {
	transactionRepository TransactionRepository
	accountRepository     AccountRepository
	generalRepository     GeneralRepository
	permissionsService    AccountPermissionsService
	logger                *logging.Logger
}

var _ GeneralRepository = &generalRepository.Repository{}
var _ TransactionRepository = &transactionRepository.TransactionRepository{}
var _ AccountPermissionsService = &accountPermissions.Service{}

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
	GetAccountPermissions(account accountModel.Account) accountPermissions.AccountPermissions
}

type AccountRepository interface {
	GetAccounts(context.Context, accountRepoModel.GetAccountsReq) ([]accountModel.Account, error)
}

// CreateTransaction создает новую транзакцию
func (s *Service) CreateTransaction(ctx context.Context, transaction transactionModel.CreateTransactionReq) (id uint32, err error) {

	// Проверяем доступ пользователя к счетам
	if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.Accounts, transaction.Necessary.UserID, []uint32{transaction.AccountFromID, transaction.AccountToID}); err != nil {
		return id, err
	}

	// Получаем счета
	_accounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{
		IDs: []uint32{transaction.AccountFromID, transaction.AccountToID},
	})
	if err != nil {
		return id, err
	}
	accountsMap := slice.ToMap(_accounts, func(account accountModel.Account) uint32 { return account.ID })

	// Получаем разрешения счетов
	permissionsAccountFrom := s.permissionsService.GetAccountPermissions(accountsMap[transaction.AccountFromID])
	permissionsAccountTo := s.permissionsService.GetAccountPermissions(accountsMap[transaction.AccountToID])

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
	return s.transactionRepository.CreateTransaction(ctx, transaction.ConvertToRepoReq())
}

func (s *Service) GetTransactions(ctx context.Context, filters transactionModel.GetTransactionsReq) (transactions []transactionModel.Transaction, err error) {

	// Проверяем доступ пользователя к группам счетов
	filters.AccountGroupIDs = s.generalRepository.GetAvailableAccountGroups(filters.Necessary.UserID)

	// Получаем все транзакции
	if transactions, err = s.transactionRepository.GetTransactions(ctx, filters); err != nil {
		return nil, err
	}

	// Заполняем массив ID транзакций
	transactionIDs := make([]uint32, len(transactions))
	for i, transaction := range transactions {
		transactionIDs[i] = transaction.ID
	}

	return transactions, nil
}

// UpdateTransaction редактирует транзакцию
func (s *Service) UpdateTransaction(ctx context.Context, fields transactionModel.UpdateTransactionReq) error {

	// Проверяем доступ пользователя к транзакции
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Transactions, fields.Necessary.UserID, []uint32{fields.ID}); err != nil {
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
		if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Accounts, fields.Necessary.UserID, accountsIDs); err != nil {
			return err
		}
	}

	// Проверяем доступ пользователя к транзакции
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Transactions, fields.Necessary.UserID, []uint32{fields.ID}); err != nil {
		return err
	}

	// Изменяем данные транзакции
	return s.transactionRepository.UpdateTransaction(ctx, fields)
}

// DeleteTransaction удаляет транзакцию
func (s *Service) DeleteTransaction(ctx context.Context, id transactionModel.DeleteTransactionReq) error {

	// Проверяем доступ пользователя к транзакции
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Transactions, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем транзакцию
	return s.transactionRepository.DeleteTransaction(ctx, id.ID, id.Necessary.UserID)
}

func New(
	transactionRepository TransactionRepository,
	accountRepository AccountRepository,
	generalRepository GeneralRepository,
	accountPermissions AccountPermissionsService,
	logger *logging.Logger,
) *Service {
	return &Service{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
		generalRepository:     generalRepository,
		permissionsService:    accountPermissions,
		logger:                logger,
	}
}

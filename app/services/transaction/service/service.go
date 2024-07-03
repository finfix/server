package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/pkg/slices"
	accountModel "server/app/services/account/model"
	"server/app/services/account/model/accountType"
	accountRepository "server/app/services/account/repository"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/accountPermissions"
	"server/app/services/generalRepository"
	"server/app/services/generalRepository/checker"
	tagModel "server/app/services/tag/model"
	tagRepository "server/app/services/tag/repository"
	transactionModel "server/app/services/transaction/model"
	"server/app/services/transaction/model/transactionType"
	transactionRepository "server/app/services/transaction/repository"
	transactionRepoModel "server/app/services/transaction/repository/model"
)

type Service struct {
	transactionRepository TransactionRepository
	accountRepository     AccountRepository
	generalRepository     GeneralRepository
	permissionsService    AccountPermissionsService
	tagRepository         TagRepository
}

var _ TransactionRepository = &transactionRepository.TransactionRepository{}
var _ AccountRepository = &accountRepository.Repository{}
var _ GeneralRepository = &generalRepository.Repository{}
var _ AccountPermissionsService = &accountPermissions.Service{}
var _ TagRepository = &tagRepository.TagRepository{}

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

type TagRepository interface {
	GetTagsToTransactions(context.Context, tagModel.GetTagsToTransactionsReq) ([]tagModel.TagToTransaction, error)
	LinkTagsToTransaction(context.Context, []uint32, uint32) error
	UnlinkTagsFromTransaction(context.Context, []uint32, uint32) error
}

// CreateTransaction создает новую транзакцию
func (s *Service) CreateTransaction(ctx context.Context, transaction transactionModel.CreateTransactionReq) (id uint32, err error) {

	// Проверяем доступ пользователя к счетам
	if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.Accounts, transaction.Necessary.UserID, []uint32{transaction.AccountFromID, transaction.AccountToID}); err != nil {
		return id, err
	}

	// Получаем счета
	_accounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{ //nolint:exhaustruct
		IDs: []uint32{transaction.AccountFromID, transaction.AccountToID},
	})
	if err != nil {
		return id, err
	}
	accountsMap := slices.ToMap(_accounts, func(account accountModel.Account) uint32 { return account.ID })

	// Проверяем, может ли пользователь использовать счета
	if err = s.transactionAndAccountTypesValidation(
		accountsMap[transaction.AccountFromID],
		accountsMap[transaction.AccountToID],
		transaction.Type,
	); err != nil {
		return id, err
	}

	// Получаем разрешения счетов
	permissionsAccountFrom := s.permissionsService.GetAccountPermissions(accountsMap[transaction.AccountFromID])
	permissionsAccountTo := s.permissionsService.GetAccountPermissions(accountsMap[transaction.AccountToID])

	// Проверяем, что счета можно использовать для создания транзакции
	if !permissionsAccountFrom.CreateTransaction || !permissionsAccountTo.CreateTransaction {
		return id, errors.BadRequest.New("Нельзя создать транзакцию для этих счетов",
			errors.ParamsOption(
				"AccountFromID", transaction.AccountFromID,
				"AccountGroupFromID", accountsMap[transaction.AccountFromID].AccountGroupID,
				"AccountToID", transaction.AccountToID,
				"AccountGroupToID", accountsMap[transaction.AccountToID].AccountGroupID,
			),
		)
	}

	// Проверяем, что счета находятся в одной группе
	if accountsMap[transaction.AccountFromID].AccountGroupID != accountsMap[transaction.AccountToID].AccountGroupID {
		return id, errors.BadRequest.New("Счета находятся в разных группах",
			errors.ParamsOption(
				"AccountFromID", transaction.AccountFromID,
				"AccountToID", transaction.AccountToID,
			))
	}

	return id, s.generalRepository.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем транзакцию
		id, err = s.transactionRepository.CreateTransaction(ctx, transaction.ConvertToRepoReq())
		if err != nil {
			return err
		}

		// Если переданы теги
		if len(transaction.TagIDs) != 0 {
			if err = s.updateTransactionTags(ctx, transaction.Necessary.UserID, id, transaction.TagIDs); err != nil {
				return err
			}
		}
		return nil
	})
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

	// Получаем транзакцию
	transactions, err := s.transactionRepository.GetTransactions(ctx, transactionModel.GetTransactionsReq{ //nolint:exhaustruct
		IDs: []uint32{fields.ID},
	})
	if err != nil {
		return err
	}
	if len(transactions) == 0 {
		return errors.NotFound.New("Транзакция не найдена",
			errors.ParamsOption("ID", fields.ID),
		)
	}
	transaction := transactions[0]

	// Если в запросе есть изменение счетов, то проверяем доступ пользователя к ним
	if fields.AccountFromID != nil || fields.AccountToID != nil {
		if fields.AccountFromID != nil {
			transaction.AccountFromID = *fields.AccountFromID
		}
		if fields.AccountToID != nil {
			transaction.AccountToID = *fields.AccountToID
		}

		// Проверяем доступ пользователя к счетам
		if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.Accounts, fields.Necessary.UserID, []uint32{transaction.AccountFromID, transaction.AccountToID}); err != nil {
			return err
		}

		// Получаем счета
		_accounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{ //nolint:exhaustruct
			IDs: []uint32{transaction.AccountFromID, transaction.AccountToID},
		})
		if err != nil {
			return err
		}
		accountsMap := slices.ToMap(_accounts, func(account accountModel.Account) uint32 { return account.ID })

		// Проверяем соответствие типов счета и типа транзакции
		if err = s.transactionAndAccountTypesValidation(
			accountsMap[transaction.AccountFromID],
			accountsMap[transaction.AccountToID],
			transaction.Type,
		); err != nil {
			return err
		}
	}

	return s.generalRepository.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Если в запросе есть изменение тегов
		if fields.TagIDs != nil {
			if err := s.updateTransactionTags(ctxTx, fields.Necessary.UserID, fields.ID, *fields.TagIDs); err != nil {
				return err
			}
		}

		// Изменяем данные транзакции
		return s.transactionRepository.UpdateTransaction(ctxTx, fields)
	})
}

func (s *Service) transactionAndAccountTypesValidation(accountFrom, accountTo accountModel.Account, tranType transactionType.Type) error {

	var accesses string
	var isAccess bool

	// Проверяем, что типы счетов выбраны правильно для этой транзакции
	switch tranType {
	case transactionType.Income:
		isAccess = accountFrom.Type == accountType.Earnings && slices.In(accountTo.Type, accountType.Regular, accountType.Debt)
		accesses = "Earnings -> [Regular, Debt]"
	case transactionType.Transfer:
		isAccess = slices.In(accountFrom.Type, accountType.Regular, accountType.Debt) && slices.In(accountTo.Type, accountType.Regular, accountType.Debt)
		accesses = "[Regular, Debt] -> [Regular, Debt]"
	case transactionType.Consumption:
		isAccess = slices.In(accountFrom.Type, accountType.Regular, accountType.Debt) && accountTo.Type == accountType.Expense
		accesses = "[Regular, Debt] -> Expense"
	case transactionType.Balancing:
		isAccess = accountFrom.Type == accountType.Balancing && slices.In(accountTo.Type, accountType.Regular, accountType.Debt)
		accesses = "Balancing -> [Regular, Debt]"
	}

	if !isAccess {
		return errors.BadRequest.New("Неверно выбраны типы счетов",
			errors.ParamsOption(
				"TransactionType", tranType,
				"AccountFromID", accountFrom.ID,
				"AccountToID", accountTo.ID,
				"Accesses", accesses,
			),
		)
	}

	return nil
}

func (s *Service) updateTransactionTags(ctx context.Context, userID, transactionID uint32, tagIDs []uint32) error {

	// Проверяем доступ пользователя к тегам
	if len(tagIDs) > 0 {
		if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Tags, userID, tagIDs); err != nil {
			return err
		}
	}

	// Получаем все теги, привязанные к транзакции
	transactionTags, err := s.tagRepository.GetTagsToTransactions(ctx, tagModel.GetTagsToTransactionsReq{ //nolint:exhaustruct
		TransactionIDs: []uint32{transactionID},
	})
	if err != nil {
		return err
	}

	existTagIDs := slices.GetFields(transactionTags, func(tag tagModel.TagToTransaction) uint32 { return tag.TagID })

	toDelete, toCreate := slices.JoinExclusive(existTagIDs, tagIDs)

	if len(toDelete) > 0 {
		// Удаляем связи тегов с транзакцией
		if err = s.tagRepository.UnlinkTagsFromTransaction(ctx, toDelete, transactionID); err != nil {
			return err
		}
	}
	if len(toCreate) > 0 {
		// Создаем связи тегов с транзакцией
		if err = s.tagRepository.LinkTagsToTransaction(ctx, toCreate, transactionID); err != nil {
			return err
		}
	}
	return nil
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
	tagRepository TagRepository,

) *Service {
	return &Service{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
		generalRepository:     generalRepository,
		permissionsService:    accountPermissions,
		tagRepository:         tagRepository,
	}
}

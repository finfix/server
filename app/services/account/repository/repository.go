package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/sql"
	"server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
)

type Repository struct {
	db     sql.SQL
	logger *logging.Logger
}

func (repo *Repository) CreateAccountGroup(ctx context.Context, req model.CreateAccountGroupReq) (uint32, error) {
	return repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.account_groups (
			  name,
              available_budget,
              currency_signatura
            ) VALUES (?, ?, ?)`,
		req.Name,
		req.AvailableBudget,
		req.Currency)
}

func (repo *Repository) GetAccountGroups(ctx context.Context, filters model.GetAccountGroupsReq) (accountGroups []model.AccountGroup, err error) {

	var (
		queryArgs = []string{"ag.id != ?"}
		args      = []any{0}
	)

	if len(filters.AccountGroupIDs) != 0 {
		_queryArgs, _args, err := repo.db.In(`ag.id IN (?)`, filters.AccountGroupIDs)
		if err != nil {
			return accountGroups, err
		}
		queryArgs = append(queryArgs, _queryArgs)
		args = append(args, _args...)
	}

	if filters.Necessary.UserID != 0 {
		queryArgs = append(queryArgs, `utag.user_id = ?`)
		args = append(args, filters.Necessary.UserID)
	}

	if len(queryArgs) == 0 {
		return accountGroups, errors.BadRequest.New("No filters")
	}

	query := fmt.Sprintf(`
			SELECT ag.*
			FROM coin.account_groups ag
    		  JOIN coin.users_to_account_groups utag ON utag.account_group_id = ag.id
			WHERE %s`, strings.Join(queryArgs, " AND "))

	// Выполняем запрос
	if err = repo.db.Select(ctx, &accountGroups, query, args...); err != nil {
		return accountGroups, err
	}

	return accountGroups, nil
}

// CreateAccount создает новый счет
func (repo *Repository) CreateAccount(ctx context.Context, account accountRepoModel.CreateAccountReq) (id uint32, serialNumber uint32, err error) {

	// Создаем счет
	id, err = repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.accounts (
			  budget_amount,
			  name,
			  icon_id,
			  type_signatura,
			  currency_signatura,
			  visible,
			  account_group_id,
			  accounting,
			  budget_gradual_filling,
			  is_parent,
			  budget_fixed_sum,
			  budget_days_offset,        
			  parent_account_id,
			  created_by_user_id,
			  time_create,
			  serial_number
		  	) SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, COALESCE(MAX(serial_number), 0) + 1 
		  	  FROM coin.accounts`,
		account.Budget.Amount,
		account.Name,
		account.IconID,
		account.Type,
		account.Currency,
		account.Visible,
		account.AccountGroupID,
		account.Accounting,
		account.Budget.GradualFilling,
		account.IsParent,
		account.Budget.FixedSum,
		account.Budget.DaysOffset,
		account.ParentAccountID,
		account.UserID,
		time.Now(),
	)
	if err != nil {
		return id, serialNumber, err
	}
	if err := repo.db.QueryRow(ctx, `
		SELECT serial_number 
		FROM coin.accounts 
		WHERE id = ?`, id).
		Scan(&serialNumber); err != nil {
		return id, serialNumber, err
	}
	return id, serialNumber, nil
}

// GetAccounts возвращает все счета, удовлетворяющие фильтрам
func (repo *Repository) GetAccounts(ctx context.Context, req accountRepoModel.GetAccountsReq) (accounts []model.Account, err error) {

	// Создаем конструктор запроса
	var (
		variables []any
		reqFields []string
	)

	if len(req.AccountGroupIDs) != 0 {
		_query, _args, err := repo.db.In(`a.account_group_id IN (?)`, req.AccountGroupIDs)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.IDs) != 0 {
		_query, _args, err := repo.db.In(`a.id IN (?)`, req.IDs)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.Types) != 0 {
		_query, _args, err := repo.db.In(`a.type_signatura IN (?)`, req.Types)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.Currencies) != 0 {
		_query, _args, err := repo.db.In(`a.currency_signatura IN (?)`, req.Currencies)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if len(req.ParentAccountIDs) != 0 {
		_query, _args, err := repo.db.In(`a.parent_account_id IN (?)`, req.ParentAccountIDs)
		if err != nil {
			return accounts, err
		}
		reqFields = append(reqFields, _query)
		variables = append(variables, _args...)
	}
	if req.IsParent != nil {
		reqFields = append(reqFields, `a.is_parent = ?`)
		variables = append(variables, req.IsParent)
	}
	if req.Accounting != nil {
		reqFields = append(reqFields, `a.accounting = ?`)
		variables = append(variables, req.Accounting)
	}
	if req.Visible != nil {
		reqFields = append(reqFields, `a.visible = ?`)
		variables = append(variables, req.Visible)
	}

	if len(reqFields) == 0 {
		return accounts, errors.BadRequest.New("No filters")
	}

	// Выполняем запрос
	query := fmt.Sprintf(`
			SELECT a.*
			FROM coin.accounts a
			WHERE %v`, strings.Join(reqFields, " AND "))
	if err = repo.db.Select(ctx, &accounts, query, variables...); err != nil {
		return accounts, err
	}

	return accounts, nil
}

// CalculateRemainderAccounts возвращает остатки счетов
func (repo *Repository) CalculateRemainderAccounts(ctx context.Context, req accountRepoModel.CalculateRemaindersAccountsReq) (map[uint32]float64, error) {

	var queryFields []string
	var args []any

	// Добавляем в запрос счета
	if len(req.IDs) != 0 {
		_query, _args, err := repo.db.In(
			`a.id IN (?)`, req.IDs,
		)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	// Добавляем в запрос типы счетов
	if len(req.Types) != 0 {
		_query, _args, err := repo.db.In(
			`a.type_signatura IN (?)`, req.Types,
		)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	// Добавляем в запрос группы счетов
	if len(req.AccountGroupIDs) != 0 {
		_query, _args, err := repo.db.In(
			`a.account_group_id IN (?)`, req.AccountGroupIDs,
		)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	// Добавляем в запрос даты
	if req.DateFrom != nil {
		queryFields = append(queryFields, `t.date_transaction >= ?`)
		args = append(args, req.DateFrom)
	}
	if req.DateTo != nil {
		queryFields = append(queryFields, `t.date_transaction < ?`)
		args = append(args, req.DateTo)
	}

	// Получаем сумму всех транзакций из счетов
	query := fmt.Sprintf(`
			SELECT t.account_to_id AS id, COALESCE(SUM(t.amount_to), 0) AS remainder
			FROM coin.transactions t
			JOIN coin.accounts a ON t.account_to_id = a.id
			WHERE %v
			GROUP BY t.account_to_id`,
		strings.Join(queryFields, " AND "))

	var amountsArray []struct {
		ID        uint32  `db:"id"`
		Remainder float64 `db:"remainder"`
	}

	// Вычисляем сумму всех транзакций из счетов, id - сумма из
	if err := repo.db.Select(ctx, &amountsArray, query, args...); err != nil {
		return nil, err
	}

	// Формируем мапу с суммой исходящих транзакций в виде map[accountID]amountTransactions
	amountFromAccount := make(map[uint32]float64, len(amountsArray))
	for _, remainder := range amountsArray {
		amountFromAccount[remainder.ID] = remainder.Remainder
	}

	// Получаем сумму всех транзакций в счета
	query = fmt.Sprintf(`
			SELECT t.account_from_id AS id, COALESCE(SUM(t.amount_from), 0) AS remainder
			FROM coin.transactions t
			JOIN coin.accounts a ON t.account_from_id = a.id
			WHERE %v
			GROUP BY t.account_from_id`,
		strings.Join(queryFields, " AND "))

	// Вычисляем сумму всех транзакций в счета, id - сумма в
	if err := repo.db.Select(ctx, &amountsArray, query, args...); err != nil {
		return nil, err
	}

	// Формируем мапу с суммой входящих транзакций в виде map[accountID]amountTransactions
	amountToAccount := make(map[uint32]float64, len(amountFromAccount))
	for _, remainder := range amountsArray {
		amountToAccount[remainder.ID] = remainder.Remainder
	}

	// Проходим по всем счетам и вычисляем остаток разницей суммы в и из счета, формируем новую мапу
	amountMapToAccountID := make(map[uint32]float64)
	for id := range amountToAccount {
		amountMapToAccountID[id] = amountFromAccount[id] - amountToAccount[id]
	}

	// Если счета нет в списке транзакций, то добавляем его со значением -сумма из счета
	for id := range amountFromAccount {
		if _, ok := amountMapToAccountID[id]; !ok {
			amountMapToAccountID[id] = amountFromAccount[id]
		}
	}

	return amountMapToAccountID, nil
}

// UpdateAccount обновляет счет
func (repo *Repository) UpdateAccount(ctx context.Context, updateReqs map[uint32]accountRepoModel.UpdateAccountReq) error {

	for id, fields := range updateReqs {

		var (
			queryFields []string
			args        []any
		)

		// Добавляем в запрос только те поля, которые необходимо обновить
		if fields.IconID != nil {
			queryFields = append(queryFields, "icon_id = ?")
			args = append(args, fields.IconID)
		}
		if fields.Accounting != nil {
			queryFields = append(queryFields, "accounting = ?")
			args = append(args, fields.Accounting)
		}
		if fields.Name != nil {
			queryFields = append(queryFields, "name = ?")
			args = append(args, fields.Name)
		}
		if fields.Visible != nil {
			queryFields = append(queryFields, "visible = ?")
			args = append(args, fields.Visible)
		}
		if fields.Budget.DaysOffset != nil {
			queryFields = append(queryFields, "budget_days_offset = ?")
			args = append(args, fields.Budget.DaysOffset)
		}
		if fields.Budget.Amount != nil {
			queryFields = append(queryFields, "budget_amount = ?")
			args = append(args, fields.Budget.Amount)
		}
		if fields.Budget.FixedSum != nil {
			queryFields = append(queryFields, "budget_fixed_sum = ?")
			args = append(args, fields.Budget.FixedSum)
		}
		if fields.Budget.GradualFilling != nil {
			queryFields = append(queryFields, "budget_gradual_filling = ?")
			args = append(args, fields.Budget.GradualFilling)
		}
		if fields.Currency != nil {
			queryFields = append(queryFields, "currency_signatura = ?")
			args = append(args, fields.Currency)
		}
		if fields.ParentAccountID != nil {
			if *fields.ParentAccountID == 0 {
				queryFields = append(queryFields, "parent_account_id = NULL")
			} else {
				queryFields = append(queryFields, "parent_account_id = ?")
				args = append(args, fields.ParentAccountID)
			}
		}

		if len(queryFields) == 0 {
			if fields.Remainder == nil {
				return errors.BadRequest.New("No fields to update")
			}
			return nil
		}

		// Конструируем запрос
		query := fmt.Sprintf(`
				UPDATE coin.accounts 
				  SET %s 
				WHERE id = ?`,
			strings.Join(queryFields, ", "),
		)
		args = append(args, id)

		// Обновляем счета
		err := repo.db.Exec(ctx, query, args...)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteAccount удаляет счет
func (repo *Repository) DeleteAccount(_ context.Context, _ uint32) error {
	panic("implement me")
}

func (repo *Repository) SwitchAccountsBetweenThemselves(ctx context.Context, id1, id2 uint32) error {
	return repo.db.Exec(ctx, `
			UPDATE coin.accounts
			SET serial_number = CASE
    		  WHEN id = ? THEN (SELECT serial_number FROM coin.accounts WHERE id = ?)
    		  WHEN id = ? THEN (SELECT serial_number FROM coin.accounts WHERE id = ?)
    		  ELSE serial_number
    		  END
			WHERE id IN (?, ?);`,
		id1, id2,
		id2, id1,
		id1, id2,
	)
}

func New(db sql.SQL, logger *logging.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

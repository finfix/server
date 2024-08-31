package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"

	"server/pkg/errors"
	"server/pkg/sql"
	"server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
)

type Repository struct {
	db sql.SQL
}

// CreateAccount создает новый счет
func (repo *Repository) CreateAccount(ctx context.Context, account accountRepoModel.CreateAccountReq) (id uint32, serialNumber uint32, err error) {

	// Получаем текущий максимальный серийный номер
	row, err := repo.db.QueryRow(ctx, `
			SELECT COALESCE(MAX(serial_number), 1) AS serial_number
			FROM coin.accounts 
			WHERE account_group_id = ?`,
		account.AccountGroupID,
	)
	if err != nil {
		return id, serialNumber, err
	}
	if err = row.Scan(&serialNumber); err != nil {
		return id, serialNumber, err
	}
	serialNumber++

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
			  accounting_in_header,
			  accounting_in_charts,
			  budget_gradual_filling,
			  is_parent,
			  budget_fixed_sum,
			  budget_days_offset,        
			  parent_account_id,
			  created_by_user_id,
			  datetime_create,
			  serial_number
		  	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		account.Budget.Amount,
		account.Name,
		account.IconID,
		account.Type,
		account.Currency,
		account.Visible,
		account.AccountGroupID,
		account.AccountingInHeader,
		account.AccountingInCharts,
		account.Budget.GradualFilling,
		account.IsParent,
		account.Budget.FixedSum,
		account.Budget.DaysOffset,
		account.ParentAccountID,
		account.UserID,
		account.DatetimeCreate,
		serialNumber,
	)
	if err != nil {
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
	if req.AccountingInHeader != nil {
		reqFields = append(reqFields, `a.accounting_in_header = ?`)
		variables = append(variables, req.AccountingInHeader)
	}
	if req.AccountingInCharts != nil {
		reqFields = append(reqFields, `a.accounting_in_charts = ?`)
		variables = append(variables, req.AccountingInCharts)
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

func (repo *Repository) ChangeSerialNumbers(ctx context.Context, accountGroupID, oldValue, newValue uint32) error {

	var req string
	args := []any{
		accountGroupID,
	}

	if newValue < oldValue {
		req = `UPDATE coin.accounts
			SET serial_number = serial_number + 1
			WHERE account_group_id = ? 
			  AND serial_number >= ?
			  AND serial_number < ?`
		args = append(args, newValue, oldValue)
	} else {
		req = `UPDATE coin.accounts
			SET serial_number = serial_number - 1
			WHERE account_group_id = ? 
			  AND serial_number > ?
			  AND serial_number <= ?`
		args = append(args, oldValue, newValue)
	}

	return repo.db.Exec(ctx, req, args...)
}

// CalculateRemainderAccounts возвращает остатки счетов
func (repo *Repository) CalculateRemainderAccounts(ctx context.Context, req accountRepoModel.CalculateRemaindersAccountsReq) (map[uint32]decimal.Decimal, error) {

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
		ID        uint32          `db:"id"`
		Remainder decimal.Decimal `db:"remainder"`
	}

	// Вычисляем сумму всех транзакций из счетов, id - сумма из
	if err := repo.db.Select(ctx, &amountsArray, query, args...); err != nil {
		return nil, err
	}

	// Формируем мапу с суммой исходящих транзакций в виде map[accountID]amountTransactions
	amountFromAccount := make(map[uint32]decimal.Decimal, len(amountsArray))
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
	amountToAccount := make(map[uint32]decimal.Decimal, len(amountFromAccount))
	for _, remainder := range amountsArray {
		amountToAccount[remainder.ID] = remainder.Remainder
	}

	// Проходим по всем счетам и вычисляем остаток разницей суммы в и из счета, формируем новую мапу
	amountMapToAccountID := make(map[uint32]decimal.Decimal)
	for id := range amountToAccount {
		amountMapToAccountID[id] = amountFromAccount[id].Sub(amountToAccount[id])
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
		if fields.AccountingInHeader != nil {
			queryFields = append(queryFields, "accounting_in_header = ?")
			args = append(args, fields.AccountingInHeader)
		}
		if fields.AccountingInCharts != nil {
			queryFields = append(queryFields, "accounting_in_charts = ?")
			args = append(args, fields.AccountingInCharts)
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
		if fields.SerialNumber != nil {
			queryFields = append(queryFields, "serial_number = ?")
			args = append(args, fields.SerialNumber)
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
				  SET %v 
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
func (repo *Repository) DeleteAccount(ctx context.Context, id uint32) error {
	return repo.db.Exec(ctx, `DELETE FROM coin.accounts WHERE id = ?`, id)
}

func New(db sql.SQL, ) *Repository {
	return &Repository{
		db: db,
	}
}

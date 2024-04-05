package repository

import (
	"context"
	"fmt"
	"strings"

	"server/app/enum/accountType"
	"server/app/enum/transactionType"
	"server/app/internal/services/account/model"
	"server/pkg/datetime/date"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/sql"
)

type Repository struct {
	db     *sql.DB
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

func (repo *Repository) CalculateBalancingAmount(ctx context.Context, accountGroupIDs []uint32, dateFrom, dateTo date.Date) (amount []model.BalancingAmount, err error) {
	return amount, repo.db.Select(ctx, &amount, `
			SELECT a.account_group_id, a.currency_signatura, SUM(t.amount_to) AS amount 
			FROM coin.transactions t
			  JOIN coin.accounts a ON a.id = t.account_to_id 
			WHERE t.type_signatura = ?
			  AND t.date_transaction >= ?
			  AND t.date_transaction < ?
			GROUP BY a.account_group_id, a.currency_signatura `,
		transactionType.Balancing,
		dateFrom,
		dateTo,
	)
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

	if filters.UserID != 0 {
		queryArgs = append(queryArgs, `utag.user_id = ?`)
		args = append(args, filters.UserID)
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

// Create создает новый счет
func (repo *Repository) Create(ctx context.Context, account model.CreateReq) (uint32, error) {

	// Создаем счет
	return repo.db.ExecWithLastInsertID(ctx, `
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
			  serial_number
		  	) SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, COALESCE(MAX(serial_number), 0) + 1 
		  	  FROM coin.accounts`,
		account.Budget.Amount,
		account.Name,
		account.IconID,
		account.Type,
		account.Currency,
		true,
		account.AccountGroupID,
		account.Accounting,
		account.Budget.GradualFilling,
		false,
		account.Budget.FixedSum,
		account.Budget.DaysOffset,
	)
}

// Get возвращает все счета, удовлетворяющие фильтрам
func (repo *Repository) Get(ctx context.Context, req model.GetReq) (accounts []model.Account, err error) {

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
	if req.Type != nil {
		reqFields = append(reqFields, `a.type_signatura = ?`)
		variables = append(variables, req.Type)
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

// CalculateExpensesAndEarnings возвращает сумму расходов и доходов за период
func (repo *Repository) CalculateExpensesAndEarnings(ctx context.Context, accountGroupIDs []uint32, dateFrom, dateTo date.Date) (map[uint32]float64, error) {

	var (
		remaindersArray []struct {
			AccountID uint32  `db:"id"`
			Remainder float64 `db:"remainders"`
		}
		args []any
	)

	questions, _args, err := repo.db.In(`?`, accountGroupIDs)
	if err != nil {
		return nil, err
	}

	args = append(args, _args...)
	args = append(args, dateFrom, dateTo)
	args = append(args, _args...)
	args = append(args, dateFrom, dateTo)

	query := fmt.Sprintf(`
			SELECT 
			  a.id AS id,
			  SUM(t.amount_from) AS remainders
			FROM coin.transactions t
			  JOIN coin.accounts a ON a.id = t.account_from_id
			WHERE a.account_group_id IN (%v)
			  AND t.date_transaction >= ?
			  AND t.date_transaction < ?
			  AND a.type_signatura = 'earnings'
			GROUP BY a.id, t.account_from_id
			
			UNION
			
			SELECT 
			  a.id,
			  SUM(t.amount_to) AS remaindersArray
			FROM coin.transactions t
			  JOIN coin.accounts a ON a.id = t.account_to_id
			WHERE a.account_group_id IN (%v)
			  AND t.date_transaction >= ?
			  AND t.date_transaction < ?
			  AND a.type_signatura = 'expense'
			GROUP BY a.id, t.account_to_id`, questions, questions)

	// Вычисляем остатки счетов earnings и expense
	if err = repo.db.Select(ctx, &remaindersArray, query, args...); err != nil {
		return nil, err
	}

	// Формируем ответ в виде map[accountID]remainders
	remaindersMap := make(map[uint32]float64, len(remaindersArray))
	for _, remainder := range remaindersArray {
		remaindersMap[remainder.AccountID] = remainder.Remainder
	}

	return remaindersMap, nil
}

// CalculateRemainderAccounts возвращает остатки счетов
func (repo *Repository) CalculateRemainderAccounts(ctx context.Context, accountGroupIDs []uint32, dateTo *date.Date) (map[uint32]float64, error) {

	var (
		queryFields = []string{`a.type_signatura NOT IN (?, ?)`}
		args        = []any{accountType.Earnings, accountType.Expense}
	)

	_query, _args, err := repo.db.In(
		`a.account_group_id IN (?)`, accountGroupIDs,
	)
	if err != nil {
		return nil, err
	}
	queryFields = append(queryFields, _query)
	args = append(args, _args...)

	// TODO: Обыграть по-другому, иначе транзакции в будущем времени не учитываются
	/*if dateTo != nil {
		queryFields = append(queryFields, `t.date_transaction < ?`)
		args = append(args, dateTo)
	}*/

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

// Update обновляет счет
func (repo *Repository) Update(ctx context.Context, fields model.UpdateReq) error {

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
	args = append(args, fields.ID)

	// Обновляем счета
	return repo.db.Exec(ctx, query, args...)
}

// GetRemainder возвращает остаток на счете
func (repo *Repository) GetRemainder(ctx context.Context, id uint32) (remainder float64, err error) {

	// Отнимаем сумму транзакций, поступивших на счет из суммы транзакций, снятых со счета
	return remainder, repo.db.QueryRow(ctx, `
			SELECT 
			  (
			    SELECT COALESCE(SUM(amount_to), 0) 
				FROM coin.transactions 
				WHERE account_to_id = ?
			  ) - (
			    SELECT COALESCE(SUM(amount_from), 0) 
				FROM coin.transactions
				WHERE account_from_id = ?
			  ) AS remainder`,
		id,
		id,
	).Scan(&remainder)
}

// Delete удаляет счет
func (repo *Repository) Delete(_ context.Context, _ uint32) error {
	panic("implement me")
}

func (repo *Repository) Switch(ctx context.Context, id1, id2 uint32) error {
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

func New(db *sql.DB, logger *logging.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

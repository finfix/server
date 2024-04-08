package generalRepository

import (
	"context"
	"fmt"
	"time"

	"server/app/services/action/model/enum"
	"server/app/services/generalRepository/checker"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/sql"
)

type Repository struct {
	db       *sql.DB
	logger   *logging.Logger
	accesses map[uint32]map[uint32]struct{}
}

// WithinTransaction принимает коллбэк, который будет выполнен в рамках транзакции
func (repo *Repository) WithinTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	// begin transaction
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	// Запускаем коллбэк
	err = callback(sql.InjectTx(ctx, tx))
	if err != nil {
		// Если произошла ошибка, откатываем изменения
		_ = tx.Rollback()
		return err
	}
	// Если ошибок нет, подтверждаем изменения
	return tx.Commit()
}

// UpdateLastSync Обновляет время последней синхронизации для устройства
func (repo *Repository) UpdateLastSync(ctx context.Context, deviceID string) error {
	return repo.db.Exec(ctx, `
			UPDATE coin.sessions 
			SET last_sync = ? 
			WHERE device_id = ?`,
		time.Now(),
		deviceID,
	)
}

// CreateAction создает новый лог действия пользователя
func (repo *Repository) CreateAction(ctx context.Context, actionType enum.ActionType, deviceID string, userID, objectID uint32) error {

	// Добавляем лог
	if err := repo.db.Exec(ctx, `
			INSERT INTO coin.action_history(
			  action_type_signatura, 
		      user_id, 
		      object_id, 
		      action_time
		    ) VALUES (?, ?, ?, ?)`,
		actionType,
		userID,
		objectID,
		time.Now(),
	); err != nil {
		return err
	}

	// Обновляем время последней синхронизации
	return repo.UpdateLastSync(ctx, deviceID)
}

// GetCurrencies возвращает список валют и их курсов
func (repo *Repository) GetCurrencies(ctx context.Context) (map[string]float64, error) {
	var currencies []struct {
		Name string  `db:"signatura"`
		Rate float64 `db:"rate"`
	}
	if err := repo.db.Select(ctx, &currencies, `
			SELECT * 
			FROM coin.currencies`,
	); err != nil {
		return nil, err
	}

	rates := make(map[string]float64, len(currencies))
	for _, currency := range currencies {
		rates[currency.Name] = currency.Rate
	}

	return rates, nil
}

// CheckAccess проверяет, имеет ли пользователь доступ к указанным идентификаторам объектов
func (repo *Repository) CheckAccess(ctx context.Context, checkType checker.CheckType, userID uint32, ids []uint32) error {

	accessedAccountGroupIDs := repo.GetAvailableAccountGroups(userID)

	if len(accessedAccountGroupIDs) == 0 {
		return errors.NotFound.New("Нет доступных объектов", errors.Options{
			Params: map[string]any{
				"UserID": userID,
				"IDs":    ids,
				"Type":   checkType,
			},
		})
	}

	var (
		countAccess              uint32
		questionsAccountGroupIDs string
		argsAGs                  []any
		questionsIDs             string
		argsIDs                  []any
		args                     []any
		errString                string
		query                    string
	)

	// Добавляем в запрос условие на проверяемые идентификаторы групп счетов
	questionsAccountGroupIDs, argsAGs, err := repo.db.In("?", accessedAccountGroupIDs)
	if err != nil {
		return err
	}

	// Добавляем в запрос условие на проверяемые идентификаторы
	questionsIDs, argsIDs, err = repo.db.In("?", ids)
	if err != nil {
		return err
	}

	// В зависимости от проверяемого типа, выбираем таблицу
	switch checkType {

	case checker.Accounts:
		query = fmt.Sprintf(`
				SELECT COUNT(*)
				FROM coin.accounts a
				WHERE a.account_group_id IN (%v)
				AND a.id IN (%v)`,
			questionsAccountGroupIDs,
			questionsIDs,
		)
		args = append(args, argsAGs...)
		args = append(args, argsIDs...)
		errString = "счетам"

	case checker.AccountGroups:
		for _, accountGroupID := range ids {
			if _, ok := repo.accesses[userID][accountGroupID]; !ok {
				return errors.Forbidden.New("Access denied", errors.Options{
					Params: map[string]any{
						"UserID": userID,
						"IDs":    ids,
						"Type":   checkType,
					},
					HumanText: fmt.Sprintf("Вы не имеете доступа к группе счетов с ID %v", accountGroupID),
				})
			}
		}
		return nil

	case checker.Transactions:
		if len(ids) != 1 {
			return errors.InternalServer.New("Невозможно проверить доступ к нескольким транзакциям")
		}
		errString = "транзакциям"
		query = fmt.Sprintf(`
				SELECT COUNT(*)
				FROM coin.transactions t 
				JOIN coin.accounts a1 ON a1.id = t.account_from_id 
				JOIN coin.accounts a2 ON a2.id = t.account_to_id
				WHERE a1.account_group_id IN (%v)
				AND a2.account_group_id IN (%v)
				AND t.id IN (%v)`,
			questionsAccountGroupIDs,
			questionsAccountGroupIDs,
			questionsIDs,
		)
		args = append(args, argsAGs...)
		args = append(args, argsAGs...)
		args = append(args, argsIDs...)
	}

	// Смотрим количество записей, которые удовлетворяют условию
	if err = repo.db.QueryRow(ctx, query, args...).Scan(&countAccess); err != nil {
		return err
	}

	// Если количество записей не равно количеству проверяемых идентификаторов, то возвращаем ошибку
	if countAccess != uint32(len(ids)) {
		return errors.Forbidden.New("Access denied", errors.Options{
			Params: map[string]any{
				"UserID": userID,
				"IDs":    ids,
				"Type":   checkType,
			},
			HumanText: fmt.Sprintf("Вы не имеете доступа к %s", errString),
		})
	}

	return nil
}

func (repo *Repository) getAccesses(ctx context.Context) (_ map[uint32]map[uint32]struct{}, err error) {

	usersToAccountsGroups := make(map[uint32]map[uint32]struct{})

	var result []struct {
		UserID         uint32 `db:"user_id"`
		AccountGroupID uint32 `db:"account_group_id"`
	}

	if err = repo.db.Select(ctx, &result, `
			SELECT u.id AS user_id, ag.id AS account_group_id
			FROM coin.account_groups ag
			JOIN coin.users_to_account_groups utag ON utag.account_group_id = ag.id 
			JOIN coin.users u ON utag.user_id = u.id`); err != nil {
		return nil, err
	}

	for _, item := range result {
		if _, ok := usersToAccountsGroups[item.UserID]; !ok {
			usersToAccountsGroups[item.UserID] = make(map[uint32]struct{})
		}
		usersToAccountsGroups[item.UserID][item.AccountGroupID] = struct{}{}
	}

	return usersToAccountsGroups, nil
}

func (repo *Repository) GetAvailableAccountGroups(userID uint32) []uint32 {
	availableAccountGroupIDs := make([]uint32, 0, len(repo.accesses[userID]))
	for accountGroupID := range repo.accesses[userID] {
		availableAccountGroupIDs = append(availableAccountGroupIDs, accountGroupID)
	}
	return availableAccountGroupIDs
}

func New(dbx *sql.DB, logger *logging.Logger) (_ *Repository, err error) {

	repository := &Repository{
		db:     dbx,
		logger: logger,
	}

	logger.Info("Получаем доступы пользователей к объектам")
	err = repository.refreshAccesses(true)
	if err != nil {
		return nil, err
	}
	go repository.refreshAccesses(false)

	return repository, nil
}

func (repo *Repository) refreshAccesses(doOnce bool) error {
	for {
		var err error
		repo.accesses, err = repo.getAccesses(context.Background())
		if doOnce {
			return err
		}
		if err != nil {
			repo.logger.Error(err)
		}

		time.Sleep(time.Minute)
	}
}

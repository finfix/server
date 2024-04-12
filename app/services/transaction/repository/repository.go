package repository

import (
	"context"
	defErr "errors"
	"fmt"
	"strings"
	"time"

	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/sql"
	"server/app/services/transaction/model"
)

// Create создает новую транзакцию
func (repo *TransactionRepository) Create(ctx context.Context, req model.CreateReq) (id uint32, err error) {

	// Создаем транзакцию
	if id, err = repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.transactions (
    		  type_signatura, 
              date_transaction, 
              account_from_id, 
              account_to_id, 
              amount_from, 
              amount_to,  
              note,  
              is_executed,  
              date_create,
			  created_by_user_id
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.Type,
		req.DateTransaction,
		req.AccountFromID,
		req.AccountToID,
		req.AmountFrom,
		req.AmountTo,
		req.Note,
		req.IsExecuted,
		time.Now(),
		req.UserID,
	); err != nil {
		return id, err
	}
	return id, nil
}

// CreateTags создает и привязывает теги к транзакции
func (repo *TransactionRepository) CreateTags(ctx context.Context, tags []string, transactionID uint32) error {

	// TODO: Оптимизировать, см. кронтаб
	// Для каждого тега делаем запись в таблицу
	for _, tag := range tags {
		if err := repo.db.Exec(ctx, `
				INSERT INTO coin.tags_to_transaction (
            	  transaction_id, 
            	  tag_id
            	) VALUES (? ,?)`,
			transactionID,
			tag,
		); err != nil {
			return err
		}
	}

	return nil
}

// Update редактирует транзакцию
func (repo *TransactionRepository) Update(ctx context.Context, fields model.UpdateReq) error {

	// Изменяем показатели транзакции
	var (
		args        []any
		queryFields []string
		query       string
	)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.IsExecuted != nil {
		queryFields = append(queryFields, `s_executed = ?`)
		args = append(args, fields.IsExecuted)
	}
	if fields.AccountFromID != nil {
		queryFields = append(queryFields, `account_from_id = ?`)
		args = append(args, fields.AccountFromID)
	}
	if fields.AccountToID != nil {
		queryFields = append(queryFields, `account_to_id = ?`)
		args = append(args, fields.AccountToID)
	}
	if fields.AmountFrom != nil {
		queryFields = append(queryFields, `amount_from = ?`)
		args = append(args, fields.AmountFrom)
	}
	if fields.AmountTo != nil {
		queryFields = append(queryFields, `amount_to = ?`)
		args = append(args, fields.AmountTo)
	}
	// TODO: Убрать, когда реализуем конвертеры
	if fields.DateTransaction != nil {
		queryFields = append(queryFields, `date_transaction = ?`)
		args = append(args, fields.DateTransaction)
	}
	if fields.Note != nil {
		if *fields.Note == "" {
			queryFields = append(queryFields, `note = NULL`)
		} else {
			queryFields = append(queryFields, `note = ?`)
			args = append(args, fields.Note)
		}
	}
	if len(queryFields) == 0 {
		return errors.BadRequest.New("No fields to update")
	}

	// Конструируем запрос
	query = fmt.Sprintf(`
 			   UPDATE coin.transactions 
               SET %v
			   WHERE id = ?`,
		strings.Join(queryFields, ", "),
	)
	args = append(args, fields.ID)

	// Создаем транзакцию
	return repo.db.Exec(ctx, query, args...)
}

// Delete удаляет транзакцию
func (repo *TransactionRepository) Delete(ctx context.Context, id, userID uint32) error {

	// Удаляем транзакцию
	rows, err := repo.db.ExecWithRowsAffected(ctx, `
			   DELETE FROM coin.transactions 
			   WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}

	// Проверяем, что в базе данных что-то изменилось
	if rows == 0 {
		return errors.NotFound.New("No such model found for model", errors.Options{
			Params: map[string]any{
				"UserID":        userID,
				"TransactionID": id,
			},
		})
	}

	return nil
}

// Get возвращает все транзакции по фильтрам
func (repo *TransactionRepository) Get(ctx context.Context, req model.GetReq) (transactions []model.Transaction, err error) {

	var (
		args        []any
		queryFields []string
	)

	_query, _args, err := repo.db.In(`account_group_id IN (?)`, req.AccountGroupIDs)
	if err != nil {
		return nil, err
	}
	queryFields = append(queryFields, fmt.Sprintf(`(a1.%v OR a2.%v)`, _query, _query))
	args = append(args, _args...)
	args = append(args, _args...)

	// Добавляем фильтры
	if req.AccountID != nil {
		queryFields = append(queryFields, `(a1.id = ? OR a2.id = ?)`)
		args = append(args, *req.AccountID, *req.AccountID)
	}
	if req.Type != nil {
		queryFields = append(queryFields, `t.type_signatura = ?`)
		args = append(args, *req.Type)
	}
	if req.DateFrom != nil {
		queryFields = append(queryFields, `t.date_transaction >= ?`)
		args = append(args, req.DateFrom)
	}
	if req.DateTo != nil {
		queryFields = append(queryFields, `t.date_transaction < ?`)
		args = append(args, req.DateTo)
	}

	// Конструируем запрос
	query := fmt.Sprintf(`SELECT t.*
		   FROM coin.transactions t
			 JOIN coin.accounts a1 ON a1.id = t.account_from_id
			 JOIN coin.accounts a2 ON a2.id = t.account_to_id
		   WHERE %v
           ORDER BY 
             t.date_transaction DESC,
             t.date_create DESC`,
		strings.Join(queryFields, " AND "),
	)

	if req.Limit != nil {
		query += ` LIMIT ?`
		args = append(args, *req.Limit)
	}
	if req.Offset != nil {
		query += ` OFFSET ?`
		args = append(args, *req.Offset)
	}

	// Получаем транзакции
	if err = repo.db.Select(ctx, &transactions, query, args...); err != nil {
		if defErr.Is(err, context.Canceled) {
			return nil, errors.ClientReject.New("HTTP connection terminated")
		}
		return nil, err
	}

	return transactions, nil
}

// GetTags возвращает все теги по списку транзакций
// TODO: Поменять на мапу
func (repo *TransactionRepository) GetTags(ctx context.Context, ids []uint32) (tags []model.Tag, err error) {

	// Конструируем запрос
	query, args, err := repo.db.In(`
			SELECT * 
			FROM coin.tags_to_transaction 
			WHERE transaction_id IN (?)`,
		ids,
	)
	if err != nil {
		return nil, err
	}

	// Получаем теги
	return tags, repo.db.Select(ctx, &tags, query, args...)
}

type TransactionRepository struct {
	db     sql.SQL
	logger *logging.Logger
}

func New(db sql.SQL, logger *logging.Logger) *TransactionRepository {
	return &TransactionRepository{
		db:     db,
		logger: logger,
	}
}

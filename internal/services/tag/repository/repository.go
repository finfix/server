package repository

import (
	"context"
	defErr "errors"
	"fmt"
	"strings"

	"server/pkg/errors"
	"server/pkg/sql"
	"server/internal/services/tag/model"
	tagRepoModel "server/internal/services/tag/repository/model"
)

type TagRepository struct {
	db sql.SQL
}

// CreateTag создает новую подкатегорию
func (repo *TagRepository) CreateTag(ctx context.Context, req tagRepoModel.CreateTagReq) (id uint32, err error) {

	// Создаем подкатегорию
	if id, err = repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.tags (
    		  name,
			  account_group_id,
			  created_by_user_id,
			  datetime_create
            ) VALUES (?, ?, ?, ?)`,
		req.Name,
		req.AccountGroupID,
		req.CreatedByUserID,
		req.DatetimeCreate,
	); err != nil {
		return id, err
	}
	return id, nil
}

// UpdateTag редактирует подкатегорию
func (repo *TagRepository) UpdateTag(ctx context.Context, fields model.UpdateTagReq) error {

	// Изменяем показатели подкатегории
	var (
		args        []any
		queryFields []string
		query       string
	)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.Name != nil {
		queryFields = append(queryFields, `name = ?`)
		args = append(args, fields.Name)
	}
	if len(queryFields) == 0 {
		return errors.BadRequest.New("No fields to update")
	}

	// Конструируем запрос
	query = fmt.Sprintf(`
 			   UPDATE coin.tags 
               SET %v
			   WHERE id = ?`,
		strings.Join(queryFields, ", "),
	)
	args = append(args, fields.ID)

	// Редактируем подкатегорию
	return repo.db.Exec(ctx, query, args...)
}

// DeleteTag удаляет подкатегорию
func (repo *TagRepository) DeleteTag(ctx context.Context, id, userID uint32) error {

	// Удаляем подкатегорию
	rows, err := repo.db.ExecWithRowsAffected(ctx, ` DELETE FROM coin.tags WHERE id = ?`, id)
	if err != nil {
		return err
	}

	// Проверяем, что в базе данных что-то изменилось
	if rows == 0 {
		return errors.NotFound.New("No such model found for model",
			errors.ParamsOption("UserID", userID, "TagID", id),
		)
	}

	return nil
}

// GetTags возвращает все подкатегории по фильтрам
func (repo *TagRepository) GetTags(ctx context.Context, req model.GetTagsReq) (tags []model.Tag, err error) {

	var (
		args        []any
		queryFields []string
	)

	_query, _args, err := repo.db.In(`account_group_id IN (?)`, req.AccountGroupIDs)
	if err != nil {
		return nil, err
	}
	queryFields = append(queryFields, _query)
	args = append(args, _args...)

	// Конструируем запрос
	query := fmt.Sprintf(`SELECT *
		   FROM coin.tags 
		   WHERE %v`,
		strings.Join(queryFields, " AND "),
	)

	// Получаем подкатегории
	if err = repo.db.Select(ctx, &tags, query, args...); err != nil {
		if defErr.Is(err, context.Canceled) {
			return nil, errors.ClientReject.New("HTTP connection terminated")
		}
		return nil, err
	}

	return tags, nil
}

// LinkTagsToTransaction привязывает подкатегории к транзакции
func (repo *TagRepository) LinkTagsToTransaction(ctx context.Context, tagIDs []uint32, transactionID uint32) error {

	queryValuesTemplate := "(?, ?)"
	queryArgs := make([]string, 0, len(tagIDs))
	args := make([]any, 0, len(tagIDs)*2) // nolint:gomnd

	for _, tagID := range tagIDs {
		queryArgs = append(queryArgs, queryValuesTemplate)
		args = append(args, transactionID, tagID)
	}

	query := fmt.Sprintf(`
		INSERT INTO coin.tags_to_transaction (
            transaction_id,
            tag_id
        ) VALUES %v`,
		strings.Join(queryArgs, ", "),
	)

	return repo.db.Exec(ctx, query, args...)
}

// UnlinkTagsFromTransaction отвязывает подкатегории от транзакции
func (repo *TagRepository) UnlinkTagsFromTransaction(ctx context.Context, tagIDs []uint32, transactionID uint32) error {

	args := make([]any, 0, len(tagIDs)+1)

	args = append(args, transactionID)

	queryIn, _args, err := repo.db.In(`tag_id IN (?)`, tagIDs)
	if err != nil {
		return err
	}
	args = append(args, _args...)

	// Удаляем связи между подкатегориями и транзакцией
	query := fmt.Sprintf(`
		DELETE FROM coin.tags_to_transaction
		WHERE transaction_id = ? AND %v`,
		queryIn,
	)

	return repo.db.Exec(ctx, query, args...)
}

// GetTagsToTransactions возвращает все связи между подкатегориями и транзакциями
func (repo *TagRepository) GetTagsToTransactions(ctx context.Context, req model.GetTagsToTransactionsReq) (res []model.TagToTransaction, err error) {

	var (
		args        []any
		queryFields []string
	)

	if len(req.AccountGroupIDs) != 0 {
		_query, _args, err := repo.db.In(`ag.id IN (?)`, req.AccountGroupIDs)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}
	if len(req.TransactionIDs) != 0 {
		_query, _args, err := repo.db.In(`ttt.transaction_id IN (?)`, req.TransactionIDs)
		if err != nil {
			return nil, err
		}
		queryFields = append(queryFields, _query)
		args = append(args, _args...)
	}

	request := fmt.Sprintf(`SELECT *
		FROM coin.tags_to_transaction ttt
		JOIN coin.tags t ON t.id = ttt.tag_id
		JOIN coin.account_groups ag ON ag.id = t.account_group_id 
		WHERE %v`,
		strings.Join(queryFields, " AND "),
	)

	return res, repo.db.Select(ctx, &res, request, args...)
}

func New(db sql.SQL, ) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

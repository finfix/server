package repository

import (
	"context"
	"fmt"
	"strings"

	"pkg/errors"

	"server/internal/services/tag/model"
)

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

package repository

import (
	"context"
	"fmt"
	"strings"

	"server/pkg/errors"
	"server/internal/services/accountGroup/model"
)

// UpdateAccountGroup обновляет группу счетов
func (repo *Repository) UpdateAccountGroup(ctx context.Context, fields model.UpdateAccountGroupReq) error {

	var (
		queryFields []string
		args        []any
	)

	// Добавляем в запрос только те поля, которые необходимо обновить
	if fields.Name != nil {
		queryFields = append(queryFields, "name = ?")
		args = append(args, fields.Name)
	}
	if fields.Currency != nil {
		queryFields = append(queryFields, "currency_signatura = ?")
		args = append(args, fields.Currency)
	}
	if fields.SerialNumber != nil {
		queryFields = append(queryFields, "serial_number = ?")
		args = append(args, fields.SerialNumber)
	}
	if fields.Visible != nil {
		queryFields = append(queryFields, "visible = ?")
		args = append(args, fields.Visible)
	}

	if len(queryFields) == 0 {
		return errors.BadRequest.New("No fields to update")
	}

	// Конструируем запрос
	query := fmt.Sprintf(`
				UPDATE coin.account_groups
				  SET %v 
				WHERE id = ?`,
		strings.Join(queryFields, ", "),
	)
	args = append(args, fields.ID)

	// Обновляем группы счетов
	err := repo.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

package repository

import (
	"context"

	accountGroupRepoModel "server/internal/services/accountGroup/repository/model"
)

// CreateAccountGroup создает новую группу счетов
func (repo *Repository) CreateAccountGroup(ctx context.Context, accountGroup accountGroupRepoModel.CreateAccountGroupReq) (id uint32, serialNumber uint32, err error) {

	// Получаем текущий максимальный серийный номер группы счетов
	row, err := repo.db.QueryRow(ctx, `
			SELECT COALESCE(MAX(serial_number), 1) AS serial_number
			FROM coin.account_groups`,
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
			INSERT INTO coin.account_groups (
			  name,
			  currency_signatura,
			  visible,
			  datetime_create,
			  serial_number,
			  created_by_user_id
		  	) VALUES (?, ?, ?, ?, ?, ?)`,
		accountGroup.Name,
		accountGroup.Currency,
		accountGroup.Visible,
		accountGroup.DatetimeCreate,
		serialNumber,
		accountGroup.UserID,
	)
	if err != nil {
		return id, serialNumber, err
	}
	return id, serialNumber, nil
}

package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	accountGroupRepoModel "server/internal/services/accountGroup/repository/model"
)

// CreateAccountGroup создает новую группу счетов
func (repo *AccountGroupRepository) CreateAccountGroup(ctx context.Context, accountGroup accountGroupRepoModel.CreateAccountGroupReq) (id uint32, serialNumber uint32, err error) {

	// Получаем текущий максимальный серийный номер группы счетов
	row, err := repo.db.QueryRow(ctx, sq.
		Select("COALESCE(MAX(serial_number), 1) AS serial_number").
		From("coin.account_groups"),
	)
	if err != nil {
		return id, serialNumber, err
	}

	// Сканируем результат
	if err = row.Scan(&serialNumber); err != nil {
		return id, serialNumber, err
	}

	// Увеличиваем серийный номер для нового элемента
	serialNumber++

	// Создаем группу счетов
	id, err = repo.db.ExecWithLastInsertID(ctx, sq.
		Insert("coin.account_groups").
		SetMap(map[string]any{
			"name":            accountGroup.Name,
			"currency":        accountGroup.Currency,
			"visible":         accountGroup.Visible,
			"datetime_create": accountGroup.DatetimeCreate,
			"serial_number":   serialNumber,
			"user_id":         accountGroup.UserID,
		}),
	)
	if err != nil {
		return id, serialNumber, err
	}

	return id, serialNumber, nil
}

package repository

import (
	"context"

	transactionRepoModel "server/app/services/transaction/repository/model"
)

// CreateFile добавляет строку с данными о файле
func (repo *TransactionRepository) CreateFile(ctx context.Context, req transactionRepoModel.CreateFileReq) (id uint32, err error) {

	// Создаем изображение
	if id, err = repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.transaction_files (
			  original_file_name,
			  file_name,
              datetime_create,
			  created_by_user_id,
			  transaction_id
            ) VALUES (?, ?, ?, ?, ?)`,
		req.TransactionID,
		req.OriginalFileName,
		req.UniqueFileName,
		req.DatetimeCreate,
		req.CreatedByUserID,
	); err != nil {
		return id, err
	}

	return id, nil
}

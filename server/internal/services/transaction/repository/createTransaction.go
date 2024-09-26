package repository

import (
	"context"

	"server/internal/services/transaction/repository/model"
)

// CreateTransaction создает новую транзакцию
func (repo *TransactionRepository) CreateTransaction(ctx context.Context, req model.CreateTransactionReq) (id uint32, err error) {

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
              datetime_create,
			  created_by_user_id,
			  accounting_in_charts
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.Type,
		req.DateTransaction,
		req.AccountFromID,
		req.AccountToID,
		req.AmountFrom,
		req.AmountTo,
		req.Note,
		req.IsExecuted,
		req.DatetimeCreate,
		req.CreatedByUserID,
		req.AccountingInCharts,
	); err != nil {
		return id, err
	}
	return id, nil
}

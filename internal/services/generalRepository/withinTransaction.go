package generalRepository

import (
	"context"
	"fmt"

	"server/pkg/sql"
)

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

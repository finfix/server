package transactor

import (
	"context"
	"fmt"

	"pkg/sql"
)

type Transactor struct {
	db sql.SQL
}

func NewTransactor(db sql.SQL) *Transactor {
	return &Transactor{
		db: db,
	}
}

// WithinTransaction принимает коллбэк, который будет выполнен в рамках транзакции
func (r *Transactor) WithinTransaction(ctx context.Context, callback func(ctx context.Context) error) error {

	// Запускаем транзакцию
	tx, err := r.db.Begin(ctx)
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

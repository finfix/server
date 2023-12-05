package model

import (
	"time"

	"core/app/enum/transactionType"
	"pkg/datetime/date"
)

type Transaction struct {
	ID              uint32               `db:"id"`               // Идентификатор транзакции
	Type            transactionType.Type `db:"type_signatura"`   // Тип транзакции
	Tags            []Tag                `db:"tags"`             // Подкатегории
	AmountFrom      float64              `db:"amount_from"`      // Сумма сделки в первой валюте
	AmountTo        float64              `db:"amount_to"`        // Сумма сделки во второй валюте
	Note            string               `db:"note"`             // Заметка сделки
	AccountFromID   uint32               `db:"account_from_id"`  // Идентификатор счета списания
	AccountToID     uint32               `db:"account_to_id"`    // Идентификатор счета пополнения
	DateTransaction date.Date            `db:"date_transaction"` // Дата транзакции (пользовательские)
	IsExecuted      bool                 `db:"is_executed"`      // Исполнена операция или нет (если нет, сделки как бы не существует)
	Accounting      bool                 `db:"accounting"`       // Учитывается ли транзакция в статистике или нет
	TimeCreate      *time.Time           `db:"time_create"`      // Дата и время создания транзакции
}

type Tag struct {
	TransactionID uint32 `db:"transaction_id"`
	TagID         uint32 `db:"tag_id"`
}

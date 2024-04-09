package model

import (
	"time"

	"server/app/pkg/datetime/date"
	"server/app/services/transaction/model/transactionType"
)

type Transaction struct {
	ID              uint32               `json:"id" db:"id" minimum:"1"`                                                             // Идентификатор транзакции
	Type            transactionType.Type `json:"type" db:"type_signatura" enums:"consumption,income,transfer"`                       // Тип транзакции
	Tags            []Tag                `json:"tags"`                                                                               // Подкатегории
	AmountFrom      float64              `json:"amountFrom" db:"amount_from" minimum:"1"`                                            // Сумма сделки в первой валюте
	AmountTo        float64              `json:"amountTo" db:"amount_to" minimum:"1"`                                                // Сумма сделки во второй валюте
	Note            string               `json:"note" db:"note"`                                                                     // Заметка сделки
	AccountFromID   uint32               `json:"accountFromID" db:"account_from_id" minimum:"1"`                                     // Идентификатор счета списания
	AccountToID     uint32               `json:"accountToID" db:"account_to_id" minimum:"1"`                                         // Идентификатор счета пополнения
	DateTransaction date.Date            `json:"dateTransaction" db:"date_transaction" format:"date" swaggertype:"primitive,string"` // Дата транзакции (пользовательские)
	IsExecuted      bool                 `json:"isExecuted" db:"is_executed"`                                                        // Исполнена операция или нет (если нет, сделки как бы не существует)
	Accounting      bool                 `json:"accounting" db:"accounting"`                                                         // Учитывается ли транзакция в статистике или нет
	TimeCreate      *time.Time           `json:"timeCreate" db:"timeCreate" format:"date-time"`                                      // Дата и время создания транзакции
}

type Tag struct {
	TransactionID uint32 `json:"transactionID" db:"transaction_id" minimum:"1"`
	TagID         uint32 `json:"tagID" db:"tagID" minimum:"1"`
}

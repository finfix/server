package model

import (
	"github.com/shopspring/decimal"

	"server/app/pkg/datetime"
	"server/app/services/transaction/model/transactionType"
)

type Transaction struct {
	ID                 uint32               `json:"id" db:"id" minimum:"1"`                                                             // Идентификатор транзакции
	Type               transactionType.Type `json:"type" db:"type_signatura" enums:"consumption,income,transfer"`                       // Тип транзакции
	AmountFrom         decimal.Decimal      `json:"amountFrom" db:"amount_from" minimum:"1"`                                            // Сумма сделки в первой валюте
	AmountTo           decimal.Decimal      `json:"amountTo" db:"amount_to" minimum:"1"`                                                // Сумма сделки во второй валюте
	Note               string               `json:"note" db:"note"`                                                                     // Заметка сделки
	AccountFromID      uint32               `json:"accountFromID" db:"account_from_id" minimum:"1"`                                     // Идентификатор счета списания
	AccountToID        uint32               `json:"accountToID" db:"account_to_id" minimum:"1"`                                         // Идентификатор счета пополнения
	DateTransaction    datetime.Date        `json:"dateTransaction" db:"date_transaction" format:"date" swaggertype:"primitive,string"` // Дата транзакции (пользовательские)
	IsExecuted         bool                 `json:"isExecuted" db:"is_executed"`                                                        // Исполнена операция или нет (если нет, сделки как бы не существует)
	AccountingInCharts bool                 `json:"accountingInCharts" db:"accounting_in_charts"`                                       // Учитывается ли транзакция в графиках или нет
	CreatedByUserID    uint32               `json:"createdByUserID" db:"created_by_user_id" minimum:"1"`                                // Идентификатор пользователя, создавшего транзакцию
	DatetimeCreate     datetime.Time        `json:"datetimeCreate" db:"datetime_create" format:"date-time"`                             // Дата и время создания транзакции
}

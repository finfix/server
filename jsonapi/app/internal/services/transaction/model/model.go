package model

import (
	"time"

	"core/app/enum/transactionType"
	"pkg/datetime/date"
)

type Transaction struct {
	ID              uint32               `json:"id" minimum:"1"`                                               // Идентификатор транзакции
	Type            transactionType.Type `json:"type" enums:"consumption,income,transfer"`                     // Тип транзакции
	Tags            []Tag                `json:"tags"`                                                         // Подкатегории
	AmountFrom      float64              `json:"amountFrom" minimum:"1"`                                       // Сумма сделки в первой валюте
	AmountTo        float64              `json:"amountTo" minimum:"1"`                                         // Сумма сделки во второй валюте
	Note            string               `json:"note"`                                                         // Заметка сделки
	AccountFromID   uint32               `json:"accountFromID" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID     uint32               `json:"accountToID" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction date.Date            `json:"dateTransaction" format:"date" swaggertype:"primitive,string"` // Дата транзакции (пользовательские)
	IsExecuted      bool                 `json:"isExecuted"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
	Accounting      bool                 `json:"accounting"`                                                   // Учитывается ли транзакция в статистике или нет
	TimeCreate      *time.Time           `json:"timeCreate" format:"date-time"`                                // Дата и время создания транзакции
}

type Tag struct {
	TransactionID uint32 `json:"transactionID" minimum:"1"`
	TagID         uint32 `json:"tagID" minimum:"1"`
}

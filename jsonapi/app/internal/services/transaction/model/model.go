package model

import (
	"time"

	"core/app/enum/transactionType"
	"pkg/datetime/date"
)

type Transaction struct {
	ID              uint32               `jsonapi:"id" minimum:"1"`                                               // Идентификатор транзакции
	Type            transactionType.Type `jsonapi:"type" enums:"consumption,income,transfer"`                     // Тип транзакции
	Tags            []Tag                `jsonapi:"tags"`                                                         // Подкатегории
	AmountFrom      float64              `jsonapi:"amountFrom" minimum:"1"`                                       // Сумма сделки в первой валюте
	AmountTo        float64              `jsonapi:"amountTo" minimum:"1"`                                         // Сумма сделки во второй валюте
	Note            string               `jsonapi:"note"`                                                         // Заметка сделки
	AccountFromID   uint32               `jsonapi:"accountFromID" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID     uint32               `jsonapi:"accountToID" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction date.Date            `jsonapi:"dateTransaction" format:"date" swaggertype:"primitive,string"` // Дата транзакции (пользовательские)
	IsExecuted      bool                 `jsonapi:"isExecuted"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
	Accounting      bool                 `jsonapi:"accounting"`                                                   // Учитывается ли транзакция в статистике или нет
	TimeCreate      *time.Time           `jsonapi:"-" format:"date-time"`                                         // Дата и время создания транзакции
}

type Tag struct {
	TransactionID uint32 `jsonapi:"transactionID" minimum:"1"`
	TagID         uint32 `jsonapi:"tagID" minimum:"1"`
}

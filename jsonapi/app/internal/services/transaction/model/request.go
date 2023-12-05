package model

import (
	"core/app/enum/transactionType"
	"pkg/datetime/date"
)

type DeleteReq struct {
	ID       uint32 `jsonapi:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
	UserID   uint32 `jsonapi:"-" validate:"required" minimum:"1"`  // Идентификатор пользователя
	DeviceID string `jsonapi:"-" validate:"required"`              // Идентификатор устройства
}

type CreateReq struct {
	Type            transactionType.Type `jsonapi:"type" validate:"required"`                                                         // Тип транзакции
	AmountFrom      float64              `jsonapi:"amountFrom" validate:"required" minimum:"1"`                                       // Сумма списания с первого счета
	AmountTo        float64              `jsonapi:"amountTo" validate:"required" minimum:"1"`                                         // Сумма пополнения второго счета (в случаях меж валютной транзакции цифры отличаются)
	Note            string               `jsonapi:"note"`                                                                             // Заметка для транзакции
	AccountFromID   uint32               `jsonapi:"accountFromID" validate:"required" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID     uint32               `jsonapi:"accountToID" validate:"required" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction date.Date            `jsonapi:"dateTransaction" validate:"required" format:"date" swaggertype:"primitive,string"` // Дата транзакции
	IsExecuted      *bool                `jsonapi:"isExecuted" validate:"required"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
	UserID          uint32               `jsonapi:"-" validate:"required" minimum:"1"`                                                // Идентификатор пользователя
	DeviceID        string               `jsonapi:"-" validate:"required"`                                                            // Идентификатор устройства
}

type UpdateReq struct {
	ID              uint32     `jsonapi:"id" validate:"required" minimum:"1"`                           // Идентификатор транзакции
	UserID          uint32     `jsonapi:"-" validate:"required" minimum:"1"`                            // Идентификатор пользователя
	AmountFrom      *float64   `jsonapi:"amountFrom" minimum:"1"`                                       // Сумма списания с первого счета
	AmountTo        *float64   `jsonapi:"amountTo" minimum:"1"`                                         // Сумма пополнения второго счета
	Note            *string    `jsonapi:"note"`                                                         // Заметка для транзакции
	AccountFromID   *uint32    `jsonapi:"accountFromID" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID     *uint32    `jsonapi:"accountToID" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction *date.Date `jsonapi:"dateTransaction" format:"date" swaggertype:"primitive,string"` // Дата транзакции
	IsExecuted      *bool      `jsonapi:"isExecuted"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
	DeviceID        string     `jsonapi:"-" validate:"required"`                                        // Идентификатор устройства
}

type GetReq struct {
	UserID    uint32                `jsonapi:"-" schema:"-" validate:"required" minimum:"1"`                            // Идентификатор пользователя
	AccountID *uint32               `jsonapi:"accountID" schema:"accountID" minimum:"1"`                                // Транзакции какого счета нас интересуют
	Type      *transactionType.Type `jsonapi:"type" schema:"type" enums:"consumption,income,transfer"`                  // Тип транзакции
	DateFrom  *date.Date            `jsonapi:"dateFrom" schema:"dateFrom" format:"date" swaggertype:"primitive,string"` // Дата, от которой начинать учитывать транзакции
	DateTo    *date.Date            `jsonapi:"dateTo" schema:"dateTo" format:"date" swaggertype:"primitive,string"`     // Дата, до которой учитывать транзакции
	DeviceID  string                `jsonapi:"-" schema:"-" validate:"required"`                                        // Идентификатор устройства
	Offset    *uint32               `jsonapi:"offset" schema:"offset" minimum:"0"`                                      // Смещение относительно начала списка для пагинации
	Limit     *uint32               `jsonapi:"limit" schema:"limit" minimum:"1"`                                        // Количество транзакций в списке для пагинации
}

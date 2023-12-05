package model

import (
	"core/app/enum/transactionType"
	"pkg/datetime/date"
)

type DeleteReq struct {
	ID     uint32 // Идентификатор транзакции
	UserID uint32 // Идентификатор пользователя
}

type CreateReq struct {
	Type            transactionType.Type // Тип транзакции
	AmountFrom      float64              // Сумма списания с первого счета
	AmountTo        float64              // Сумма пополнения второго счета (в случаях меж валютной транзакции цифры отличаются)
	Note            string               // Заметка для транзакции
	AccountFromID   uint32               // Идентификатор счета списания
	AccountToID     uint32               // Идентификатор счета пополнения
	DateTransaction date.Date            // Дата транзакции
	IsExecuted      *bool                // Исполнена операция или нет (если нет, сделки как бы не существует)
	UserID          uint32               // Идентификатор пользователя
}

type UpdateReq struct {
	ID              uint32     // Идентификатор транзакции
	UserID          uint32     // Идентификатор пользователя
	AmountFrom      *float64   // Сумма списания с первого счета
	AmountTo        *float64   // Сумма пополнения второго счета
	Note            *string    // Заметка для транзакции
	AccountFromID   *uint32    // Идентификатор счета списания
	AccountToID     *uint32    // Идентификатор счета пополнения
	DateTransaction *date.Date // Дата транзакции
	IsExecuted      *bool      // Исполнена операция или нет (если нет, сделки как бы не существует)
}

type GetReq struct {
	UserID          uint32                // Идентификатор пользователя
	AccountGroupIDs []uint32              // Идентификаторы групп счетов
	AccountID       *uint32               // Транзакции какого счета нас интересуют
	Type            *transactionType.Type // Тип транзакции
	DateFrom        *date.Date            // Дата, от которой начинать учитывать транзакции
	DateTo          *date.Date            // Дата, до которой учитывать транзакции
	Offset          *uint32               // Смещение относительно начала списка для пагинации
	Limit           *uint32               // Количество транзакций в списке для пагинации
}

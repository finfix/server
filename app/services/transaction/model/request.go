package model

import (
	"server/app/pkg/datetime/date"
	"server/app/services"
	"server/app/services/transaction/model/transactionType"
	repoModel "server/app/services/transaction/repository/model"
)

type DeleteTransactionReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

type CreateTransactionReq struct {
	Necessary       services.NecessaryUserInformation
	Type            transactionType.Type `json:"type" validate:"required"`                                                         // Тип транзакции
	AmountFrom      float64              `json:"amountFrom" validate:"required" minimum:"1"`                                       // Сумма списания с первого счета
	AmountTo        float64              `json:"amountTo" validate:"required" minimum:"1"`                                         // Сумма пополнения второго счета (в случаях меж валютной транзакции цифры отличаются)
	Note            string               `json:"note"`                                                                             // Заметка для транзакции
	AccountFromID   uint32               `json:"accountFromID" validate:"required" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID     uint32               `json:"accountToID" validate:"required" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction date.Date            `json:"dateTransaction" validate:"required" format:"date" swaggertype:"primitive,string"` // Дата транзакции
	IsExecuted      *bool                `json:"isExecuted" validate:"required"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
}

func (s *CreateTransactionReq) ConvertToRepoReq() repoModel.CreateTransactionReq {
	return repoModel.CreateTransactionReq{
		Type:            s.Type,
		AmountFrom:      s.AmountFrom,
		AmountTo:        s.AmountTo,
		Note:            s.Note,
		AccountFromID:   s.AccountFromID,
		AccountToID:     s.AccountToID,
		DateTransaction: s.DateTransaction,
		IsExecuted:      s.IsExecuted,
		CreatedByUserID: s.Necessary.UserID,
	}
}

type UpdateTransactionReq struct {
	Necessary       services.NecessaryUserInformation
	ID              uint32     `json:"id" validate:"required" minimum:"1"`                           // Идентификатор транзакции
	AmountFrom      *float64   `json:"amountFrom" minimum:"1"`                                       // Сумма списания с первого счета
	AmountTo        *float64   `json:"amountTo" minimum:"1"`                                         // Сумма пополнения второго счета
	Note            *string    `json:"note"`                                                         // Заметка для транзакции
	AccountFromID   *uint32    `json:"accountFromID" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID     *uint32    `json:"accountToID" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction *date.Date `json:"dateTransaction" format:"date" swaggertype:"primitive,string"` // Дата транзакции
	IsExecuted      *bool      `json:"isExecuted"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
}

type GetTransactionsReq struct {
	Necessary       services.NecessaryUserInformation
	AccountID       *uint32               `json:"accountID" schema:"accountID" minimum:"1"`                                // Транзакции какого счета нас интересуют
	Type            *transactionType.Type `json:"type" schema:"type" enums:"consumption,income,transfer"`                  // Тип транзакции
	DateFrom        *date.Date            `json:"dateFrom" schema:"dateFrom" format:"date" swaggertype:"primitive,string"` // Дата, от которой начинать учитывать транзакции
	DateTo          *date.Date            `json:"dateTo" schema:"dateTo" format:"date" swaggertype:"primitive,string"`     // Дата, до которой учитывать транзакции
	Offset          *uint32               `json:"offset" schema:"offset" minimum:"0"`                                      // Смещение относительно начала списка для пагинации
	Limit           *uint32               `json:"limit" schema:"limit" minimum:"1"`                                        // Количество транзакций в списке для пагинации
	AccountGroupIDs []uint32              // Идентификаторы групп счетов
}

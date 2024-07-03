package model

import (
	"github.com/shopspring/decimal"

	"server/app/pkg/datetime"
	"server/app/pkg/errors"
	"server/app/services"
	"server/app/services/transaction/model/transactionType"
	repoModel "server/app/services/transaction/repository/model"
)

type DeleteTransactionReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

func (s DeleteTransactionReq) SetNecessary(information services.NecessaryUserInformation) any {
	s.Necessary = information
	return s
}

type CreateTransactionReq struct {
	Necessary          services.NecessaryUserInformation
	Type               transactionType.Type `json:"type" validate:"required"`                                                         // Тип транзакции
	AmountFrom         decimal.Decimal      `json:"amountFrom" validate:"required" minimum:"1"`                                       // Сумма списания с первого счета
	AmountTo           decimal.Decimal      `json:"amountTo" validate:"required" minimum:"1"`                                         // Сумма пополнения второго счета (в случаях меж валютной транзакции цифры отличаются)
	Note               string               `json:"note"`                                                                             // Заметка для транзакции
	AccountFromID      uint32               `json:"accountFromID" validate:"required" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID        uint32               `json:"accountToID" validate:"required" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction    datetime.Date        `json:"dateTransaction" validate:"required" format:"date" swaggertype:"primitive,string"` // Дата транзакции
	IsExecuted         *bool                `json:"isExecuted" validate:"required"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
	TagIDs             []uint32             `json:"tagIDs"`                                                                           // Идентификаторы тегов
	DatetimeCreate     datetime.Time        `json:"datetimeCreate" validate:"required"`                                               // Дата создания транзакции
	AccountingInCharts *bool                `json:"accountingInCharts" validate:"required"`                                           // Учитывается ли транзакция в графиках или нет
}

func (s CreateTransactionReq) Validate() error {
	// Валидируем поля
	if err := s.Type.Validate(); err != nil {
		return err
	}
	if s.AmountFrom.LessThanOrEqual(decimal.Zero) || s.AmountTo.LessThanOrEqual(decimal.Zero) {
		return errors.BadRequest.New("amountFrom and amountTo must be greater than 0")
	}
	return nil
}

func (s CreateTransactionReq) SetNecessary(information services.NecessaryUserInformation) any {
	s.Necessary = information
	return s
}

func (s *CreateTransactionReq) ConvertToRepoReq() repoModel.CreateTransactionReq {
	return repoModel.CreateTransactionReq{
		Type:               s.Type,
		AmountFrom:         s.AmountFrom,
		AmountTo:           s.AmountTo,
		Note:               s.Note,
		AccountFromID:      s.AccountFromID,
		AccountToID:        s.AccountToID,
		DateTransaction:    s.DateTransaction,
		IsExecuted:         *s.IsExecuted,
		CreatedByUserID:    s.Necessary.UserID,
		DatetimeCreate:     s.DatetimeCreate.Time,
		AccountingInCharts: *s.AccountingInCharts,
	}
}

type UpdateTransactionReq struct {
	Necessary          services.NecessaryUserInformation
	ID                 uint32           `json:"id" validate:"required" minimum:"1"`                           // Идентификатор транзакции
	AmountFrom         *decimal.Decimal `json:"amountFrom" minimum:"1"`                                       // Сумма списания с первого счета
	AmountTo           *decimal.Decimal `json:"amountTo" minimum:"1"`                                         // Сумма пополнения второго счета
	Note               *string          `json:"note"`                                                         // Заметка для транзакции
	AccountFromID      *uint32          `json:"accountFromID" minimum:"1"`                                    // Идентификатор счета списания
	AccountToID        *uint32          `json:"accountToID" minimum:"1"`                                      // Идентификатор счета пополнения
	DateTransaction    *datetime.Date   `json:"dateTransaction" format:"date" swaggertype:"primitive,string"` // Дата транзакции
	IsExecuted         *bool            `json:"isExecuted"`                                                   // Исполнена операция или нет (если нет, сделки как бы не существует)
	TagIDs             *[]uint32        `json:"tagIDs"`                                                       // Идентификаторы тегов
	AccountingInCharts *bool            `json:"accountingInCharts"`                                           // Учитывается ли транзакция в графиках или нет
}

func (s UpdateTransactionReq) Validate() error {
	if s.AmountFrom != nil && s.AmountFrom.LessThanOrEqual(decimal.Zero) {
		return errors.BadRequest.New("amountFrom must be greater than 0")
	}
	if s.AmountTo != nil && s.AmountTo.LessThanOrEqual(decimal.Zero) {
		return errors.BadRequest.New("amountTo must be greater than 0")
	}
	return nil
}

func (s UpdateTransactionReq) SetNecessary(information services.NecessaryUserInformation) any {
	s.Necessary = information
	return s
}

type GetTransactionsReq struct {
	Necessary       services.NecessaryUserInformation
	IDs             []uint32              `json:"-"`                                                                       // Идентификаторы транзакций
	AccountID       *uint32               `json:"accountID" schema:"accountID" minimum:"1"`                                // Транзакции какого счета нас интересуют
	Type            *transactionType.Type `json:"type" schema:"type" enums:"consumption,income,transfer"`                  // Тип транзакции
	DateFrom        *datetime.Date        `json:"dateFrom" schema:"dateFrom" format:"date" swaggertype:"primitive,string"` // Дата, от которой начинать учитывать транзакции
	DateTo          *datetime.Date        `json:"dateTo" schema:"dateTo" format:"date" swaggertype:"primitive,string"`     // Дата, до которой учитывать транзакции
	Offset          *uint32               `json:"offset" schema:"offset" minimum:"0"`                                      // Смещение относительно начала списка для пагинации
	Limit           *uint32               `json:"limit" schema:"limit" minimum:"1"`                                        // Количество транзакций в списке для пагинации
	AccountGroupIDs []uint32              // Идентификаторы групп счетов
}

func (s GetTransactionsReq) Validate() error {
	if err := s.Type.Validate(); err != nil {
		return err
	}
	if s.DateFrom != nil && s.DateTo != nil {
		if s.DateFrom.After(s.DateTo.Time) || s.DateFrom.Equal(s.DateTo.Time) {
			return errors.BadRequest.New("date_from must be less than date_to")
		}
	}
	return nil
}

func (s GetTransactionsReq) SetNecessary(information services.NecessaryUserInformation) any {
	s.Necessary = information
	return s
}

package model

import (
	"github.com/shopspring/decimal"

	"pkg/datetime"
	"pkg/errors"
	"pkg/necessary"

	"server/internal/services/transaction/model/transactionType"
	"server/internal/services/transaction/repository/model"
)

type CreateTransactionReq struct {
	Necessary          necessary.NecessaryUserInformation
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

func (s *CreateTransactionReq) ConvertToRepoReq() model.CreateTransactionReq {
	return model.CreateTransactionReq{
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

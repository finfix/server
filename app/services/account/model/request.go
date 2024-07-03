package model

import (
	"time"

	"github.com/shopspring/decimal"

	"server/app/pkg/datetime"
	"server/app/services"
	"server/app/services/account/model/accountType"
	repoModel "server/app/services/account/repository/model"
)

type GetAccountsReq struct {
	Necessary          services.NecessaryUserInformation
	Type               *accountType.Type `json:"type" schema:"type" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	AccountingInHeader *bool             `json:"accountingInHeader" schema:"accountingInHeader"`                              // Учитывать ли счет в шапке
	AccountingInCharts *bool             `json:"accountingInCharts" schema:"accountingInCharts"`                              // Учитывать ли счет в графиках
	AccountGroupIDs    []uint32          `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"`                        // Идентификаторы групп счетов
	DateFrom           *datetime.Date    `json:"dateFrom" schema:"dateFrom" format:"date" swaggertype:"primitive,string"`     // Дата начала выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	DateTo             *datetime.Date    `json:"dateTo" schema:"dateTo" format:"date" swaggertype:"primitive,string"`         // Дата конца выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	Visible            *bool             `json:"visible" schema:"visible"`                                                    // Видимость счета
	Currency           *string           `json:"-" schema:"-"`                                                                // Валюта счета
	IsParent           *bool             `json:"-" schema:"-"`                                                                // Является ли счет родительским
	IDs                []uint32          `json:"-" schema:"-"`
}

func (s GetAccountsReq) Validate() error {
	return s.Type.Validate()
}

// TODO: Переписать
func (s *GetAccountsReq) ConvertToRepoReq() repoModel.GetAccountsReq {
	var req repoModel.GetAccountsReq
	req.IDs = s.IDs
	req.AccountGroupIDs = s.AccountGroupIDs
	if s.Type != nil {
		req.Types = []accountType.Type{*s.Type}
	}
	req.AccountingInHeader = s.AccountingInHeader
	req.AccountingInCharts = s.AccountingInCharts
	req.Visible = s.Visible
	if s.Currency != nil {
		req.Currencies = []string{*s.Currency}
	}
	req.IsParent = s.IsParent

	return req
}

type CreateAccountReq struct {
	Necessary          services.NecessaryUserInformation
	Name               string                 `json:"name" validate:"required"`                                                          // Название счета
	IconID             uint32                 `json:"iconID" validate:"required" minimum:"1"`                                            // Идентификатор иконки
	Type               accountType.Type       `json:"type" validate:"required" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Currency           string                 `json:"currency" validate:"required"`                                                      // Валюта счета
	AccountGroupID     uint32                 `json:"accountGroupID" validate:"required" minimum:"1"`                                    // Группа счета
	AccountingInHeader *bool                  `json:"accountingInHeader" validate:"required"`                                            // Подсчет суммы счета в статистике
	AccountingInCharts *bool                  `json:"accountingInCharts" validate:"required"`                                            // Учитывать ли счет в графиках
	DatetimeCreate     datetime.Time          `json:"datetimeCreate" validate:"required"`                                                // Дата создания счета
	Remainder          decimal.Decimal        `json:"remainder"`                                                                         // Остаток средств на счету
	Budget             CreateAccountBudgetReq `json:"budget"`                                                                            // Бюджет
	IsParent           *bool                  `json:"isParent"`                                                                          // Является ли счет родительским
	ParentAccountID    *uint32                `json:"parentAccountID"`                                                                   // Идентификатор родительского счета
	Visible            *bool                  `json:"-"`                                                                                 // Видимость счета
}

func (s CreateAccountReq) Validate() error {
	return s.Type.Validate()
}

func (s CreateAccountReq) ContertToAccount() Account {
	return Account{
		ID:                 0,
		Remainder:          s.Remainder,
		Name:               s.Name,
		IconID:             s.IconID,
		Type:               s.Type,
		Currency:           s.Currency,
		Visible:            true,
		AccountGroupID:     s.AccountGroupID,
		AccountingInHeader: *s.AccountingInHeader,
		ParentAccountID:    s.ParentAccountID,
		SerialNumber:       0,
		IsParent:           *s.IsParent,
		CreatedByUserID:    s.Necessary.UserID,
		DatetimeCreate:     datetime.Time{Time: time.Now()},
		AccountingInCharts: *s.AccountingInCharts,
		AccountBudget: AccountBudget{
			Amount:         s.Budget.Amount,
			FixedSum:       s.Budget.FixedSum,
			DaysOffset:     s.Budget.DaysOffset,
			GradualFilling: *s.Budget.GradualFilling,
		},
	}
}

// TODO: Переписать
func (s CreateAccountReq) ConvertToRepoReq() repoModel.CreateAccountReq {
	return repoModel.CreateAccountReq{
		Name:               s.Name,
		IconID:             s.IconID,
		Type:               s.Type,
		Currency:           s.Currency,
		AccountGroupID:     s.AccountGroupID,
		AccountingInHeader: *s.AccountingInHeader,
		AccountingInCharts: *s.AccountingInCharts,
		Budget:             s.Budget.ConvertToRepoReq(),
		IsParent:           *s.IsParent,
		Visible:            true,
		ParentAccountID:    s.ParentAccountID,
		UserID:             s.Necessary.UserID,
		DatetimeCreate:     s.DatetimeCreate.Time,
	}
}

type CreateAccountBudgetReq struct {
	Amount         decimal.Decimal `json:"amount"`                             // Сумма
	FixedSum       decimal.Decimal `json:"fixedSum"`                           // Фиксированная сумма
	DaysOffset     uint32          `json:"daysOffset"`                         // Смещение в днях
	GradualFilling *bool           `json:"gradualFilling" validate:"required"` // Постепенное пополнение
}

// TODO: Переписать
func (s *CreateAccountBudgetReq) ConvertToRepoReq() repoModel.CreateReqBudget {
	return repoModel.CreateReqBudget{
		Amount:         s.Amount,
		FixedSum:       s.FixedSum,
		DaysOffset:     s.DaysOffset,
		GradualFilling: *s.GradualFilling,
	}
}

type UpdateAccountReq struct {
	Necessary          services.NecessaryUserInformation
	ID                 uint32                 `json:"id" validate:"required" minimum:"1"` // Идентификатор счета
	Remainder          *decimal.Decimal       `json:"remainder"`                          // Остаток средств на счету
	Name               *string                `json:"name"`                               // Название счета
	IconID             *uint32                `json:"iconID" minimum:"1"`                 // Идентификатор иконки
	Visible            *bool                  `json:"visible"`                            // Видимость счета
	AccountingInHeader *bool                  `json:"accountingInHeader"`                 // Будет ли счет учитываться в статистике
	AccountingInCharts *bool                  `json:"accountingInCharts"`                 // Будет ли счет учитываться в графиках
	Currency           *string                `json:"currencyCode"`                       // Валюта счета
	SerialNumber       *uint32                `json:"serialNumber"`                       // Порядковый номер счета
	ParentAccountID    *uint32                `json:"parentAccountID"`                    // Идентификатор родительского счета
	Budget             UpdateAccountBudgetReq `json:"budget"`                             // Месячный бюджет
}

func (s *UpdateAccountReq) ConvertToRepoReq() repoModel.UpdateAccountReq {
	return repoModel.UpdateAccountReq{
		Remainder:          s.Remainder,
		Name:               s.Name,
		IconID:             s.IconID,
		Visible:            s.Visible,
		AccountingInHeader: s.AccountingInHeader,
		AccountingInCharts: s.AccountingInCharts,
		Currency:           s.Currency,
		ParentAccountID:    s.ParentAccountID,
		Budget:             s.Budget.ConvertToRepoReq(),
		SerialNumber:       s.SerialNumber,
	}
}

type UpdateAccountBudgetReq struct {
	Amount         *decimal.Decimal `json:"amount"`         // Сумма
	FixedSum       *decimal.Decimal `json:"fixedSum"`       // Фиксированная сумма
	DaysOffset     *uint32          `json:"daysOffset"`     // Смещение в днях
	GradualFilling *bool            `json:"gradualFilling"` // Постепенное пополнение
}

func (s *UpdateAccountBudgetReq) ConvertToRepoReq() repoModel.UpdateAccountBudgetReq {
	return repoModel.UpdateAccountBudgetReq{
		Amount:         s.Amount,
		FixedSum:       s.FixedSum,
		DaysOffset:     s.DaysOffset,
		GradualFilling: s.GradualFilling,
	}
}

type DeleteAccountReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" schema:"id" validate:"required" minimum:"1"` // Идентификатор счета
}

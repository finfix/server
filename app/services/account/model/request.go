package model

import (
	"server/app/pkg/datetime/date"
	"server/app/services"
	"server/app/services/account/model/accountType"
	repoModel "server/app/services/account/repository/model"
)

type GetReq struct {
	Necessary       services.NecessaryUserInformation
	Type            *accountType.Type `json:"type" schema:"type" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Accounting      *bool             `json:"accounting" schema:"accounting"`                                              // Видимость счета
	AccountGroupIDs []uint32          `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"`                        // Идентификаторы групп счетов
	DateFrom        *date.Date        `json:"dateFrom" schema:"dateFrom" format:"date" swaggertype:"primitive,string"`     // Дата начала выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	DateTo          *date.Date        `json:"dateTo" schema:"dateTo" format:"date" swaggertype:"primitive,string"`         // Дата конца выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	Visible         *bool             `json:"visible" schema:"visible"`                                                    // Видимость счета
	Currency        *string           `json:"-" schema:"-"`                                                                // Валюта счета
	IsParent        *bool             `json:"-" schema:"-"`                                                                // Является ли счет родительским
	IDs             []uint32          `json:"-" schema:"-"`
}

// TODO: Переписать
func (s *GetReq) ConvertToRepoReq() repoModel.GetReq {
	var req repoModel.GetReq
	req.IDs = s.IDs
	req.AccountGroupIDs = s.AccountGroupIDs
	if s.Type != nil {
		req.Types = []accountType.Type{*s.Type}
	}
	req.Accounting = s.Accounting
	req.Visible = s.Visible
	if s.Currency != nil {
		req.Currencies = []string{*s.Currency}
	}
	req.IsParent = s.IsParent

	return req
}

type CreateReq struct {
	Necessary      services.NecessaryUserInformation
	Name           string           `json:"name" validate:"required"`                                                          // Название счета
	IconID         uint32           `json:"iconID" validate:"required" minimum:"1"`                                            // Идентификатор иконки
	Type           accountType.Type `json:"type" validate:"required" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Currency       string           `json:"currency" validate:"required"`                                                      // Валюта счета
	AccountGroupID uint32           `json:"accountGroupID" validate:"required" minimum:"1"`                                    // Группа счета
	Accounting     *bool            `json:"accounting" validate:"required"`                                                    // Подсчет суммы счета в статистике
	Remainder      float64          `json:"remainder"`                                                                         // Остаток средств на счету
	Budget         CreateBudgetReq  `json:"budget"`                                                                            // Бюджет
	IsParent       *bool            `json:"isParent"`                                                                          // Является ли счет родительским
	Visible        *bool            `json:"-"`                                                                                 // Видимость счета
}

// TODO: Переписать
func (s *CreateReq) ConvertToRepoReq() repoModel.CreateReq {
	return repoModel.CreateReq{
		Name:           s.Name,
		IconID:         s.IconID,
		Type:           s.Type,
		Currency:       s.Currency,
		AccountGroupID: s.AccountGroupID,
		Accounting:     *s.Accounting,
		Budget:         s.Budget.ConvertToCreateBudgetReqRepo(),
		IsParent:       *s.IsParent,
		Visible:        true,
		UserID:         s.Necessary.UserID,
	}
}

type CreateBudgetReq struct {
	Amount         float64 `json:"amount"`                             // Сумма
	FixedSum       float64 `json:"fixedSum"`                           // Фиксированная сумма
	DaysOffset     uint32  `json:"daysOffset"`                         // Смещение в днях
	GradualFilling *bool   `json:"gradualFilling" validate:"required"` // Постепенное пополнение
}

// TODO: Переписать
func (s *CreateBudgetReq) ConvertToCreateBudgetReqRepo() repoModel.CreateReqBudget {
	return repoModel.CreateReqBudget{
		Amount:         s.Amount,
		FixedSum:       s.FixedSum,
		DaysOffset:     s.DaysOffset,
		GradualFilling: *s.GradualFilling,
	}
}

type UpdateReq struct {
	Necessary       services.NecessaryUserInformation
	ID              uint32          `json:"id" validate:"required" minimum:"1"` // Идентификатор счета
	Remainder       *float64        `json:"remainder"`                          // Остаток средств на счету
	Name            *string         `json:"name"`                               // Название счета
	IconID          *uint32         `json:"iconID" minimum:"1"`                 // Идентификатор иконки
	Visible         *bool           `json:"visible"`                            // Видимость счета
	Accounting      *bool           `json:"accounting"`                         // Будет ли счет учитываться в статистике
	Currency        *string         `json:"currencyCode"`                       // Валюта счета
	ParentAccountID *uint32         `json:"parentAccountID"`                    // Идентификатор родительского счета
	Budget          UpdateBudgetReq `json:"budget"`                             // Месячный бюджет
}

func (s *UpdateReq) ConvertToRepoReq() repoModel.UpdateReq {
	var req repoModel.UpdateReq
	req.Remainder = s.Remainder
	req.Name = s.Name
	req.IconID = s.IconID
	req.Visible = s.Visible
	req.Accounting = s.Accounting
	req.Currency = s.Currency
	req.ParentAccountID = s.ParentAccountID
	req.Budget = s.Budget.ConvertToRepoReq()

	return req
}

type UpdateBudgetReq struct {
	Amount         *float64 `json:"amount"`         // Сумма
	FixedSum       *float64 `json:"fixedSum"`       // Фиксированная сумма
	DaysOffset     *uint32  `json:"daysOffset"`     // Смещение в днях
	GradualFilling *bool    `json:"gradualFilling"` // Постепенное пополнение
}

func (s *UpdateBudgetReq) ConvertToRepoReq() repoModel.UpdateBudgetReq {
	return repoModel.UpdateBudgetReq{
		Amount:         s.Amount,
		FixedSum:       s.FixedSum,
		DaysOffset:     s.DaysOffset,
		GradualFilling: s.GradualFilling,
	}
}

type DeleteReq struct {
	Necessary services.NecessaryUserInformation
	ID        uint32 `json:"id" schema:"id" validate:"required" minimum:"1"` // Идентификатор счета
}

type SwitchReq struct {
	Necessary services.NecessaryUserInformation
	ID1       uint32 `json:"id1" validate:"required" minimum:"1"` // Идентификатор первого счета
	ID2       uint32 `json:"id2" validate:"required" minimum:"1"` // Идентификатор второго счета
}

type GetAccountGroupsReq struct {
	Necessary       services.NecessaryUserInformation
	AccountGroupIDs []uint32 `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"` // Идентификаторы групп счетов
}

type CreateAccountGroupReq struct {
	Name            string  // Название группы счетов
	AvailableBudget float64 // Доступный бюджет
	Currency        string  // Валюта группы счетов
}

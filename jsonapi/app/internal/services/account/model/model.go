package model

import (
	"core/app/enum/accountType"
)

type Account struct {
	ID                   uint32           `jsonapi:"id"`                                                            // Идентификатор счета
	Budget               float64          `jsonapi:"budget"`                                                        // Месячный бюджет
	Remainder            float64          `jsonapi:"remainder"`                                                     // Остаток средств на счету
	Name                 string           `jsonapi:"name"`                                                          // Название счета
	IconID               uint32           `jsonapi:"iconID"`                                                        // Идентификатор иконки
	Type                 accountType.Type `jsonapi:"type" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Currency             string           `jsonapi:"currency"`                                                      // Валюта счета
	Visible              bool             `jsonapi:"visible"`                                                       // Видимость счета
	AccountGroupID       uint32           `jsonapi:"accountGroupID"`                                                // Идентификатор группы счета
	Accounting           bool             `jsonapi:"accounting"`                                                    // Будет ли счет учитываться в статистике
	ParentAccountID      *uint32          `jsonapi:"parentAccountID" validate:"required"`                           // Идентификатор родительского аккаунта
	GradualBudgetFilling bool             `jsonapi:"gradualBudgetFilling"`                                          // Постепенное пополнение бюджета
	SerialNumber         uint32           `jsonapi:"serialNumber"`                                                  // Порядковый номер счета
}

type AccountGroup struct {
	ID           uint32 `jsonapi:"id"`           // Идентификатор группы счетов
	Name         string `jsonapi:"name"`         // Название группы счетов
	Currency     string `jsonapi:"currency"`     // Валюта группы счетов
	SerialNumber uint32 `jsonapi:"serialNumber"` // Порядковый номер группы счетов
	Visible      bool   `jsonapi:"visible"`      // Видимость группы счетов
}

type QuickStatistic struct {
	AccountGroupID uint32  `jsonapi:"accountGroupID"` // Идентификатор группы счетов
	Currency       string  `jsonapi:"currency"`       // Валюта
	TotalRemainder float64 `jsonapi:"totalRemainder"` // Общий баланс видимых счетов
	TotalExpense   float64 `jsonapi:"totalExpense"`   // Общая сумма расходов
	TotalBudget    float64 `jsonapi:"totalBudget"`    // Общая сумма бюджетов
}

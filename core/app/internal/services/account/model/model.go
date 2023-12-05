package model

import "core/app/enum/accountType"

type Account struct {
	ID                   uint32           `db:"id" `                    // Идентификатор счета
	Budget               float64          `db:"budget"`                 // Месячный бюджет
	Remainder            float64          `db:"remainder"`              // Остаток средств на счету
	Name                 string           `db:"name"`                   // Название счета
	IconID               uint32           `db:"icon_id"`                // Идентификатор иконки
	Type                 accountType.Type `db:"type_signatura"`         // Тип счета
	Currency             string           `db:"currency_signatura"`     // Валюта счета
	Visible              bool             `db:"visible"`                // Видимость счета
	AccountGroupID       uint32           `db:"account_group_id"`       // Идентификатор группы счета
	Accounting           bool             `db:"accounting"`             // Будет ли счет учитываться в статистике
	ParentAccountID      *uint32          `db:"parent_account_id"`      // Идентификатор родительского аккаунта
	GradualBudgetFilling bool             `db:"gradual_budget_filling"` // Постепенное пополнение бюджета
	SerialNumber         uint32           `db:"serial_number"`          // Порядковый номер счета
	IsParent             bool             `db:"is_parent"`              // Является ли счет родительским
}

type AccountGroup struct {
	ID           uint32 `db:"id"`                 // Идентификатор группы счетов
	Name         string `db:"name"`               // Название группы счетов
	Currency     string `db:"currency_signatura"` // Валюта группы счетов
	SerialNumber uint32 `db:"serial_number"`      // Порядковый номер группы счетов
	Visible      bool   `db:"visible"`            // Видимость группы счетов
}

type BalancingAmount struct {
	Amount         float64 `db:"amount"`
	Currency       string  `db:"currency_signatura"`
	AccountGroupID uint32  `db:"account_group_id"`
}

type QuickStatistic struct {
	AccountGroupID uint32  // Идентификатор группы счетов
	Currency       string  // Валюта
	TotalRemainder float64 // Общий баланс видимых счетов
	TotalExpense   float64 // Общая сумма расходов
	TotalBudget    float64 // Общая сумма бюджетов
}

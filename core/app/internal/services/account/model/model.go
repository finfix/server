package model

import "core/app/enum/accountType"

type Account struct {
	ID              uint32           `db:"id" `                // Идентификатор счета
	Remainder       float64          `db:"remainder"`          // Остаток средств на счету
	Name            string           `db:"name"`               // Название счета
	IconID          uint32           `db:"icon_id"`            // Идентификатор иконки
	Type            accountType.Type `db:"type_signatura"`     // Тип счета
	Currency        string           `db:"currency_signatura"` // Валюта счета
	Visible         bool             `db:"visible"`            // Видимость счета
	AccountGroupID  uint32           `db:"account_group_id"`   // Идентификатор группы счета
	Accounting      bool             `db:"accounting"`         // Будет ли счет учитываться в статистике
	ParentAccountID *uint32          `db:"parent_account_id"`  // Идентификатор родительского аккаунта
	SerialNumber    uint32           `db:"serial_number"`      // Порядковый номер счета
	IsParent        bool             `db:"is_parent"`          // Является ли счет родительским
	Budget
}

type Budget struct {
	Amount         float64 `db:"budget_amount"`          // Сумма бюджета
	FixedSum       float64 `db:"budget_fixed_sum"`       // Фиксированная сумма
	DaysOffset     uint32  `db:"budget_days_offset"`     // Смещение в днях
	GradualFilling bool    `db:"budget_gradual_filling"` // Постепенное пополнение
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

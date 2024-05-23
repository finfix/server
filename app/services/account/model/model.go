package model

import (
	"github.com/shopspring/decimal"

	"server/app/pkg/datetime"
	"server/app/services/account/model/accountType"
)

type Account struct {
	ID                 uint32           `json:"id" db:"id"`                                                     // Идентификатор счета
	Remainder          decimal.Decimal  `json:"remainder" db:"remainder"`                                       // Остаток средств на счету
	Name               string           `json:"name" db:"name"`                                                 // Название счета
	IconID             uint32           `json:"iconID" db:"icon_id"`                                            // Идентификатор иконки
	Type               accountType.Type `json:"type" db:"type_signatura" enums:"regular,expense,debt,earnings"` // Тип счета
	Currency           string           `json:"currency" db:"currency_signatura"`                               // Валюта счета
	Visible            bool             `json:"visible" db:"visible"`                                           // Видимость счета
	AccountGroupID     uint32           `json:"accountGroupID" db:"account_group_id"`                           // Идентификатор группы счета
	AccountingInHeader bool             `json:"accountingInHeader" db:"accounting_in_header"`                   // Будет ли счет учитываться в шапке
	ParentAccountID    *uint32          `json:"parentAccountID" db:"parent_account_id" validate:"required"`     // Идентификатор родительского аккаунта
	SerialNumber       uint32           `json:"serialNumber" db:"serial_number"`                                // Порядковый номер счета
	IsParent           bool             `json:"isParent" db:"is_parent"`                                        // Является ли счет родительским
	CreatedByUserID    uint32           `json:"createdByUserID" db:"created_by_user_id"`                        // Идентификатор пользователя, создавшего счет
	DatetimeCreate     datetime.Time    `json:"datetimeCreate" db:"datetime_create"`                            // Дата и время создания счета
	AccountingInCharts bool             `json:"accountingInCharts" db:"accounting_in_charts"`                   // Учитывать ли счет в графиках
	AccountBudget      `json:"budget"`                                                                          // Бюджет
}

type AccountBudget struct {
	Amount         decimal.Decimal `json:"amount" db:"budget_amount"`                  // Сумма бюджета
	FixedSum       decimal.Decimal `json:"fixedSum" db:"budget_fixed_sum"`             // Фиксированная сумма
	DaysOffset     uint32          `json:"daysOffset" db:"budget_days_offset"`         // Смещение в днях
	GradualFilling bool            `json:"gradualFilling" db:"budget_gradual_filling"` // Постепенное пополнение
}

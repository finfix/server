package model

import (
	"time"

	"server/app/services/account/model/accountType"
)

type Account struct {
	ID              uint32           `json:"id" db:"id"`                                                     // Идентификатор счета
	Remainder       float64          `json:"remainder" db:"remainder"`                                       // Остаток средств на счету
	Name            string           `json:"name" db:"name"`                                                 // Название счета
	IconID          uint32           `json:"iconID" db:"icon_id"`                                            // Идентификатор иконки
	Type            accountType.Type `json:"type" db:"type_signatura" enums:"regular,expense,debt,earnings"` // Тип счета
	Currency        string           `json:"currency" db:"currency_signatura"`                               // Валюта счета
	Visible         bool             `json:"visible" db:"visible"`                                           // Видимость счета
	AccountGroupID  uint32           `json:"accountGroupID" db:"account_group_id"`                           // Идентификатор группы счета
	Accounting      bool             `json:"accounting" db:"accounting"`                                     // Будет ли счет учитываться в статистике
	ParentAccountID *uint32          `json:"parentAccountID" db:"parent_account_id" validate:"required"`     // Идентификатор родительского аккаунта
	SerialNumber    uint32           `json:"serialNumber" db:"serial_number"`                                // Порядковый номер счета
	IsParent        bool             `json:"isParent" db:"is_parent"`                                        // Является ли счет родительским
	CreatedByUserID *uint32          `json:"createdByUserID" db:"created_by_user_id"`                        // Идентификатор пользователя, создавшего счет
	DatetimeCreate  time.Time        `json:"datetimeCreate" db:"datetime_create"`                            // Дата и время создания счета
	AccountBudget   `json:"budget"`                                                                          // Бюджет
}

type AccountBudget struct {
	Amount         float64 `json:"amount" db:"budget_amount"`                  // Сумма бюджета
	FixedSum       float64 `json:"fixedSum" db:"budget_fixed_sum"`             // Фиксированная сумма
	DaysOffset     uint32  `json:"daysOffset" db:"budget_days_offset"`         // Смещение в днях
	GradualFilling bool    `json:"gradualFilling" db:"budget_gradual_filling"` // Постепенное пополнение
}

type AccountGroup struct {
	ID           uint32 `json:"id" db:"id"`                       // Идентификатор группы счетов
	Name         string `json:"name" db:"name"`                   // Название группы счетов
	Currency     string `json:"currency" db:"currency_signatura"` // Валюта группы счетов
	SerialNumber uint32 `json:"serialNumber" db:"serial_number"`  // Порядковый номер группы счетов
	Visible      bool   `json:"visible" db:"visible"`             // Видимость группы счетов
}
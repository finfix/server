package model

import (
	"core/app/enum/accountType"
)

type Account struct {
	ID              uint32           `json:"id"`                                                            // Идентификатор счета
	Remainder       float64          `json:"remainder"`                                                     // Остаток средств на счету
	Name            string           `json:"name"`                                                          // Название счета
	IconID          uint32           `json:"iconID"`                                                        // Идентификатор иконки
	Type            accountType.Type `json:"type" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Currency        string           `json:"currency"`                                                      // Валюта счета
	Visible         bool             `json:"visible"`                                                       // Видимость счета
	AccountGroupID  uint32           `json:"accountGroupID"`                                                // Идентификатор группы счета
	Accounting      bool             `json:"accounting"`                                                    // Будет ли счет учитываться в статистике
	ParentAccountID *uint32          `json:"parentAccountID" validate:"required"`                           // Идентификатор родительского аккаунта
	SerialNumber    uint32           `json:"serialNumber"`                                                  // Порядковый номер счета
	IsParent        bool             `json:"isParent"`                                                      // Является ли счет родительским
	Budget          Budget           `json:"budget"`                                                        // Бюджет
}

type Budget struct {
	Amount         float64 `json:"amount"`         // Сумма бюджета
	FixedSum       float64 `json:"fixedSum"`       // Фиксированная сумма
	DaysOffset     uint32  `json:"daysOffset"`     // Смещение в днях
	GradualFilling bool    `json:"gradualFilling"` // Постепенное пополнение
}

type AccountGroup struct {
	ID           uint32 `json:"id"`           // Идентификатор группы счетов
	Name         string `json:"name"`         // Название группы счетов
	Currency     string `json:"currency"`     // Валюта группы счетов
	SerialNumber uint32 `json:"serialNumber"` // Порядковый номер группы счетов
	Visible      bool   `json:"visible"`      // Видимость группы счетов
}

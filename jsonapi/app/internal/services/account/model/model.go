package model

import (
	"core/app/enum/accountType"
)

type Account struct {
	ID                   uint32           `json:"id"`                                                            // Идентификатор счета
	Budget               float64          `json:"budget"`                                                        // Месячный бюджет
	Remainder            float64          `json:"remainder"`                                                     // Остаток средств на счету
	Name                 string           `json:"name"`                                                          // Название счета
	IconID               uint32           `json:"iconID"`                                                        // Идентификатор иконки
	Type                 accountType.Type `json:"type" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Currency             string           `json:"currency"`                                                      // Валюта счета
	Visible              bool             `json:"visible"`                                                       // Видимость счета
	AccountGroupID       uint32           `json:"accountGroupID"`                                                // Идентификатор группы счета
	Accounting           bool             `json:"accounting"`                                                    // Будет ли счет учитываться в статистике
	ParentAccountID      *uint32          `json:"parentAccountID" validate:"required"`                           // Идентификатор родительского аккаунта
	GradualBudgetFilling bool             `json:"gradualBudgetFilling"`                                          // Постепенное пополнение бюджета
	SerialNumber         uint32           `json:"serialNumber"`                                                  // Порядковый номер счета
	IsParent             bool             `json:"isParent"`                                                      // Является ли счет родительским
}

type AccountGroup struct {
	ID           uint32 `json:"id"`           // Идентификатор группы счетов
	Name         string `json:"name"`         // Название группы счетов
	Currency     string `json:"currency"`     // Валюта группы счетов
	SerialNumber uint32 `json:"serialNumber"` // Порядковый номер группы счетов
	Visible      bool   `json:"visible"`      // Видимость группы счетов
}

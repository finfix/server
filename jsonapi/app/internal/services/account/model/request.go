package model

import (
	"core/app/enum/accountType"
	"pkg/datetime/date"
)

type GetReq struct {
	UserID          uint32            `json:"-" schema:"-" validate:"required"  minimum:"1"`                               // Идентификатор пользователя
	DeviceID        string            `json:"-" schema:"-" validate:"required"`                                            // Идентификатор устройства
	Type            *accountType.Type `json:"type" schema:"type" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Accounting      *bool             `json:"accounting" schema:"accounting"`                                              // Видимость счета
	AccountGroupIDs []uint32          `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"`                        // Идентификаторы групп счетов
	DateFrom        *date.Date        `json:"dateFrom" schema:"dateFrom" format:"date" swaggertype:"primitive,string"`     // Дата начала выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	DateTo          *date.Date        `json:"dateTo" schema:"dateTo" format:"date" swaggertype:"primitive,string"`         // Дата конца выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	Visible         *bool             `json:"visible" schema:"visible"`                                                    // Видимость счета
}

type CreateReq struct {
	Name                 string           `json:"name" validate:"required"`                                                          // Название счета
	IconID               uint32           `json:"iconID" validate:"required" minimum:"1"`                                            // Идентификатор иконки
	Type                 accountType.Type `json:"type" validate:"required" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Currency             string           `json:"currency" validate:"required"`                                                      // Валюта счета
	AccountGroupID       uint32           `json:"accountGroupID" validate:"required" minimum:"1"`                                    // Группа счета
	Accounting           *bool            `json:"accounting" validate:"required"`                                                    // Подсчет суммы счета в статистике
	UserID               uint32           `json:"-" validate:"required" minimum:"1"`                                                 // Идентификатор пользователя
	DeviceID             string           `json:"-" validate:"required"`                                                             // Идентификатор устройства
	Budget               float64          `json:"budget"`                                                                            // Месячный бюджет
	Remainder            float64          `json:"remainder"`                                                                         // Остаток средств на счету
	GradualBudgetFilling *bool            `json:"gradualBudgetFilling" validate:"required"`                                          // Постепенное пополнение бюджета
}

type UpdateReq struct {
	UserID               uint32   `json:"-" validate:"required" minimum:"1"`  // Идентификатор пользователя
	ID                   uint32   `json:"id" validate:"required" minimum:"1"` // Идентификатор счета
	Budget               *int32   `json:"budget" minimum:"0"`                 // Месячный бюджет
	Remainder            *float64 `json:"remainder"`                          // Остаток средств на счету
	Name                 *string  `json:"name"`                               // Название счета
	IconID               *uint32  `json:"iconID" minimum:"1"`                 // Идентификатор иконки
	Visible              *bool    `json:"visible"`                            // Видимость счета
	Accounting           *bool    `json:"accounting"`                         // Будет ли счет учитываться в статистике
	GradualBudgetFilling *bool    `json:"gradualBudgetFilling"`               // Постепенное пополнение бюджета
	DeviceID             string   `json:"-" validate:"required"`              // Идентификатор устройства
}

type DeleteReq struct {
	ID       uint32 `json:"id" schema:"id" validate:"required" minimum:"1"` // Идентификатор счета
	UserID   uint32 `json:"-" validate:"required" minimum:"1"`              // Идентификатор пользователя
	DeviceID string `json:"-" validate:"required"`                          // Идентификатор устройства
}

type SwitchReq struct {
	ID1      uint32 `json:"id_1" validate:"required" minimum:"1"` // Идентификатор первого счета
	ID2      uint32 `json:"id_2" validate:"required" minimum:"1"` // Идентификатор второго счета
	UserID   uint32 `json:"-" validate:"required"`                // Идентификатор пользователя
	DeviceID string `json:"-" validate:"required"`                // Идентификатор устройства
}

type QuickStatisticReq struct {
	UserID   uint32 `json:"-" validate:"required" minimum:"1"` // Идентификатор пользователя
	DeviceID string `json:"-" validate:"required"`             // Идентификатор устройства
}

type GetAccountGroupsReq struct {
	UserID          uint32   `json:"-" validate:"required" minimum:"1"`                    // Идентификатор пользователя
	DeviceID        string   `json:"-" validate:"required"`                                // Идентификатор устройства
	AccountGroupIDs []uint32 `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"` // Идентификаторы групп счетов
}

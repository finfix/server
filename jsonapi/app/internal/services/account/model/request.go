package model

import (
	"core/app/enum/accountType"
	"pkg/datetime/date"
)

type GetReq struct {
	UserID          uint32            `jsonapi:"-" schema:"-" validate:"required"  minimum:"1"`                               // Идентификатор пользователя
	DeviceID        string            `jsonapi:"-" schema:"-" validate:"required"`                                            // Идентификатор устройства
	Type            *accountType.Type `jsonapi:"type" schema:"type" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Accounting      *bool             `jsonapi:"accounting" schema:"accounting"`                                              // Видимость счета
	AccountGroupIDs []uint32          `jsonapi:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"`                        // Идентификаторы групп счетов
	DateFrom        *date.Date        `jsonapi:"dateFrom" schema:"dateFrom" format:"date" swaggertype:"primitive,string"`     // Дата начала выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	DateTo          *date.Date        `jsonapi:"dateTo" schema:"dateTo" format:"date" swaggertype:"primitive,string"`         // Дата конца выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	Visible         *bool             `jsonapi:"visible" schema:"visible"`                                                    // Видимость счета
}

type CreateReq struct {
	Name                 string           `jsonapi:"name" validate:"required"`                                                          // Название счета
	IconID               uint32           `jsonapi:"iconID" validate:"required" minimum:"1"`                                            // Идентификатор иконки
	Type                 accountType.Type `jsonapi:"type" validate:"required" enums:"regular,expense,credit,debt,earnings,investments"` // Тип счета
	Currency             string           `jsonapi:"currency" validate:"required"`                                                      // Валюта счета
	AccountGroupID       uint32           `jsonapi:"accountGroupID" validate:"required" minimum:"1"`                                    // Группа счета
	Accounting           *bool            `jsonapi:"accounting" validate:"required"`                                                    // Подсчет суммы счета в статистике
	UserID               uint32           `jsonapi:"-" validate:"required" minimum:"1"`                                                 // Идентификатор пользователя
	DeviceID             string           `jsonapi:"-" validate:"required"`                                                             // Идентификатор устройства
	Budget               float64          `jsonapi:"budget"`                                                                            // Месячный бюджет
	Remainder            float64          `jsonapi:"remainder"`                                                                         // Остаток средств на счету
	GradualBudgetFilling *bool            `jsonapi:"gradualBudgetFilling" validate:"required"`                                          // Постепенное пополнение бюджета
}

type UpdateReq struct {
	UserID               uint32   `jsonapi:"-" validate:"required" minimum:"1"`  // Идентификатор пользователя
	ID                   uint32   `jsonapi:"id" validate:"required" minimum:"1"` // Идентификатор счета
	Budget               *int32   `jsonapi:"budget" minimum:"0"`                 // Месячный бюджет
	Remainder            *float64 `jsonapi:"remainder"`                          // Остаток средств на счету
	Name                 *string  `jsonapi:"name"`                               // Название счета
	IconID               *uint32  `jsonapi:"iconID" minimum:"1"`                 // Идентификатор иконки
	Visible              *bool    `jsonapi:"visible"`                            // Видимость счета
	AccountGroupID       *uint32  `jsonapi:"accountGroupID" minimum:"1"`         // Идентификатор группы счета
	Accounting           *bool    `jsonapi:"accounting"`                         // Будет ли счет учитываться в статистике
	GradualBudgetFilling *bool    `jsonapi:"gradualBudgetFilling"`               // Постепенное пополнение бюджета
	DeviceID             string   `jsonapi:"-" validate:"required"`              // Идентификатор устройства
}

type DeleteReq struct {
	ID       uint32 `jsonapi:"id" schema:"id" validate:"required" minimum:"1"` // Идентификатор счета
	UserID   uint32 `jsonapi:"-" validate:"required" minimum:"1"`              // Идентификатор пользователя
	DeviceID string `jsonapi:"-" validate:"required"`                          // Идентификатор устройства
}

type SwitchReq struct {
	ID1      uint32 `jsonapi:"id_1" validate:"required" minimum:"1"` // Идентификатор первого счета
	ID2      uint32 `jsonapi:"id_2" validate:"required" minimum:"1"` // Идентификатор второго счета
	UserID   uint32 `jsonapi:"-" validate:"required"`                // Идентификатор пользователя
	DeviceID string `jsonapi:"-" validate:"required"`                // Идентификатор устройства
}

type QuickStatisticReq struct {
	UserID   uint32 `jsonapi:"-" validate:"required" minimum:"1"` // Идентификатор пользователя
	DeviceID string `jsonapi:"-" validate:"required"`             // Идентификатор устройства
}

type GetAccountGroupsReq struct {
	UserID          uint32   `jsonapi:"-" validate:"required" minimum:"1"`                    // Идентификатор пользователя
	DeviceID        string   `jsonapi:"-" validate:"required"`                                // Идентификатор устройства
	AccountGroupIDs []uint32 `jsonapi:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"` // Идентификаторы групп счетов
}

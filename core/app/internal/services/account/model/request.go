package model

import (
	"core/app/enum/accountType"
	"pkg/datetime/date"
)

type GetReq struct {
	IDs             []uint32          // Идентификаторы счетов
	UserID          uint32            // Идентификатор пользователя
	DeviceID        string            // Идентификатор устройства
	AccountGroupIDs []uint32          // Идентификатор группы счета
	Type            *accountType.Type // Тип счета
	Accounting      *bool             // Будет ли счет учитываться в статистике
	Visible         *bool             // Видимость счета
	DateFrom        *date.Date        // Дата начала выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	DateTo          *date.Date        // Дата конца выборки (Обязательна при type = expense or earnings и отсутствующем периоде)
	IsRemoveParents bool              // Удалять ли родительский счет
}

type CreateReq struct {
	Name                 string           // Название счета
	IconID               uint32           // Идентификатор иконки
	Type                 accountType.Type // Тип счета
	Currency             string           // Валюта счета
	AccountGroupID       uint32           // Группа счета
	Accounting           bool             // Подсчет суммы счета в статистике
	UserID               uint32           // Идентификатор пользователя
	Budget               float64          // Месячный бюджет
	Remainder            float64          // Остаток средств на счету
	GradualBudgetFilling bool             // Постепенное пополнение бюджета
}

type UpdateReq struct {
	UserID               uint32   // Идентификатор пользователя
	ID                   uint32   // Идентификатор счета
	Budget               *int32   // Месячный бюджет
	Remainder            *float64 // Остаток средств на счету
	Name                 *string  // Название счета
	IconID               *uint32  // Идентификатор иконки
	Visible              *bool    // Видимость счета
	Accounting           *bool    // Будет ли счет учитываться в статистике
	GradualBudgetFilling *bool    // Постепенное пополнение бюджета
}

type DeleteReq struct {
	ID     uint32 // Идентификатор счета
	UserID uint32 // Идентификатор пользователя
}

type SwitchReq struct {
	ID1    uint32 // Идентификатор первого счета
	ID2    uint32 // Идентификатор второго счета
	UserID uint32 // Идентификатор пользователя
}

type QuickStatisticReq struct {
	UserID uint32 // Идентификатор пользователя
}

type GetAccountGroupsReq struct {
	UserID          uint32   // Идентификатор пользователя
	AccountGroupIDs []uint32 // Идентификаторы групп счетов
}

type CreateAccountGroupReq struct {
	Name            string  // Название группы счетов
	AvailableBudget float64 // Доступный бюджет
	Currency        string  // Валюта группы счетов
}

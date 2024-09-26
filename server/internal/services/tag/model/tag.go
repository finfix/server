package model

import (
	"pkg/datetime"
)

type Tag struct {
	ID             uint32        `json:"id" db:"id" minimum:"1"`               // Идентификатор подкатегории
	AccountGroupID uint32        `json:"accountGroupID" db:"account_group_id"` // Идентификатор группы счетов
	Name           string        `json:"name" db:"name"`                       // Название подкатегории
	DatetimeCreate datetime.Time `json:"datetimeCreate" db:"datetime_create"`  // Дата и время создания
}

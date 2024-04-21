package model

import "time"

type CreateTagReq struct {
	CreatedByUserID uint32    // Идентификатор пользователя создавшего транзакцию
	AccountGroupID  uint32    // Идентификатор группы счетов
	Name            string    // Название подкатегории
	DatetimeCreate  time.Time // Дата и время создания
}

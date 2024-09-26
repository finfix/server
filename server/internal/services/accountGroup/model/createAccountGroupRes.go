package model

type CreateAccountGroupRes struct {
	ID           uint32 `json:"id"`           // Идентификатор созданной группы счетов
	SerialNumber uint32 `json:"serialNumber"` // Порядковый номер группы счетов
}

package model

type GetRes struct {
	Accounts []Account
}

type CreateRes struct {
	ID           uint32 `json:"id"`           // Идентификатор созданного счета
	SerialNumber uint32 `json:"serialNumber"` // Порядковый номер счета
}

type GetAccountGroupsRes struct {
	AccountGroups []AccountGroup
}

package model

type GetRes struct {
	Accounts []Account
}

type CreateRes struct {
	ID                     uint32 `json:"id"`                               // Идентификатор созданного счета
	SerialNumber           uint32 `json:"serialNumber"`                     // Порядковый номер счета
	BalancingTransactionID uint32 `json:"balancingTransactionID,omitempty"` // Идентификатор транзакции балансировки
}

type UpdateRes struct {
	BalancingTransactionID       uint32 `json:"balancingTransactionID,omitempty"`       // Идентификатор транзакции
	BalancingAccountID           uint32 `json:"balancingAccountID,omitempty"`           // Идентификатор балансировочного счета
	BalancingAccountSerialNumber uint32 `json:"balancingAccountSerialNumber,omitempty"` // Порядковый номер балансировочного счета
}

type GetAccountGroupsRes struct {
	AccountGroups []AccountGroup
}

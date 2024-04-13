package model

type GetRes struct {
	Accounts []Account
}

type CreateAccountRes struct {
	ID                     uint32  `json:"id"`                                         // Идентификатор созданного счета
	SerialNumber           uint32  `json:"serialNumber"`                               // Порядковый номер счета
	BalancingTransactionID *uint32 `json:"balancingTransactionID" validate:"required"` // Идентификатор транзакции балансировки
}

type UpdateAccountRes struct {
	BalancingTransactionID       *uint32 `json:"balancingTransactionID" validate:"required"`       // Идентификатор транзакции
	BalancingAccountID           *uint32 `json:"balancingAccountID" validate:"required"`           // Идентификатор балансировочного счета
	BalancingAccountSerialNumber *uint32 `json:"balancingAccountSerialNumber" validate:"required"` // Порядковый номер балансировочного счета
}

type GetAccountGroupsRes struct {
	AccountGroups []AccountGroup
}

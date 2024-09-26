package model

type UpdateAccountRes struct {
	BalancingAccountID           *uint32 `json:"balancingAccountID" validate:"required"`           // Идентификатор балансировочного счета
	BalancingTransactionID       *uint32 `json:"balancingTransactionID" validate:"required"`       // Идентификатор транзакции
	BalancingAccountSerialNumber *uint32 `json:"balancingAccountSerialNumber" validate:"required"` // Порядковый номер балансировочного счета
}

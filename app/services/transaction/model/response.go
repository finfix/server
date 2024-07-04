package model

type CreateTransactionRes struct {
	ID uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

type GetTransactionsRes struct {
	Transactions []Transaction `json:"transactions"`
}

type CreateFileRes struct {
	ID uint32 `json:"id"`
}

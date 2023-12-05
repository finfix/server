package model

type CreateRes struct {
	ID uint32 // Идентификатор транзакции
}

type GetRes struct {
	Transactions []Transaction
}

package model

type CreateRes struct {
	ID uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

type GetRes struct {
	Transactions []Transaction `json:"transactions"`
}

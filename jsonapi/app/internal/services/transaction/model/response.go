package model

type CreateRes struct {
	ID uint32 `jsonapi:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

type GetRes struct {
	Transactions []Transaction `jsonapi:"transactions"`
}

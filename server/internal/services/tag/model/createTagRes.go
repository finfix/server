package model

type CreateTagRes struct {
	ID uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

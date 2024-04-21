package model

type CreateTagRes struct {
	ID uint32 `json:"id" validate:"required" minimum:"1"` // Идентификатор транзакции
}

type TagToTransaction struct {
	TagID         uint32 `json:"tagID" minimum:"1" db:"tag_id"`                  // Идентификатор подкатегории
	TransactionID uint32 `json:"transactionID"  minimum:"1" db:"transaction_id"` // Идентификатор транзакции
}

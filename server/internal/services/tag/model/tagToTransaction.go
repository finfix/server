package model

type TagToTransaction struct {
	TagID         uint32 `json:"tagID" minimum:"1" db:"tag_id"`                  // Идентификатор подкатегории
	TransactionID uint32 `json:"transactionID"  minimum:"1" db:"transaction_id"` // Идентификатор транзакции
}

package model

type AuthRes struct {
	Tokens `json:"token"`     // Токены доступа
	ID     uint32 `json:"id"` // Идентификатор пользователя
}

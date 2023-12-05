package model

type AuthRes struct {
	Token `json:"token"`     // Токены доступа
	ID    uint32 `json:"id"` // Идентификатор пользователя
}

type RefreshTokensRes struct {
	Token `json:"token"` // Токены доступа
}

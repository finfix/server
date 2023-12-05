package model

type AuthRes struct {
	Token `jsonapi:"token"`     // Токены доступа
	ID    uint32 `jsonapi:"id"` // Идентификатор пользователя
}

type RefreshTokensRes struct {
	Token `jsonapi:"token"` // Токены доступа
}

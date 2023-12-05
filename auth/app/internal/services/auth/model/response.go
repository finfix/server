package model

type AuthRes struct {
	Token        // Токены доступа
	ID    uint32 // Идентификатор пользователя
}

type RefreshTokensRes struct {
	Token // Токены доступа
}

package model

type Token struct {
	AccessToken  string `jsonapi:"accessToken"`  // Токен доступа
	RefreshToken string `jsonapi:"refreshToken"` // Токен восстановления доступа
}

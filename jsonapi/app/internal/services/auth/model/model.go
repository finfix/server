package model

type Token struct {
	AccessToken  string `json:"accessToken"`  // Токен доступа
	RefreshToken string `json:"refreshToken"` // Токен восстановления доступа
}

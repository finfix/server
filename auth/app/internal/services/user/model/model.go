package model

import "time"

type User struct {
	ID                    uint32    // Идентификатор пользователя
	Name                  string    // Имя пользователя
	Email                 string    // Электронная почта
	PasswordHash          string    // Хэш пароля
	VerificationEmailCode *string   // Временный код, приходящий на почту
	TimeCreate            time.Time // Дата создания аккаунта
}
